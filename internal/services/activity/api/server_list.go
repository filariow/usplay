package api

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) List(ctx context.Context, req *activitycomm.ListActivitiesRequest) (*activitycomm.ListActivitiesReply, error) {
	ids := make([]uuid.UUID, len(req.FilterIds))
	for idx, i := range req.FilterIds {
		id, err := uuid.Parse(i)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "ID is not valid: %v", err)
		}
		ids[idx] = id
	}

	acts, err := s.repo.List(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of activities: %v", err)
	}

	if len(acts) == 0 {
		return &activitycomm.ListActivitiesReply{Activities: nil}, nil
	}

	// collecting the needed ActivityType details
	atrChan, ordChan := s.getActivities(ctx, acts), s.getOrders(ctx, acts)
	orders, actTypes := <-ordChan, <-atrChan

	// building response with optional data
	activities := toActivityComm(acts, orders, actTypes)

	return &activitycomm.ListActivitiesReply{
		Activities: activities,
	}, nil
}

func toActivityComm(acts storage.Activities, orders map[string]*ordercomm.Order,
	actTypes map[string]*activitytypecomm.ActivityType) []*activitycomm.Activity {
	activities := make([]*activitycomm.Activity, len(acts))
	for idx, act := range acts {
		// if info about the ActivityType was not retrieved, add the info you have
		actType, ok := actTypes[act.ActivityTypeID]
		if !ok {
			actType = &activitytypecomm.ActivityType{Id: act.ActivityTypeID}
		}

		order, ok := orders[act.OrderID]
		if !ok {
			order = &ordercomm.Order{Id: act.OrderID}
		}

		from, _ := ptypes.TimestampProto(act.Period.From)
		to, _ := ptypes.TimestampProto(act.Period.To)
		activities[idx] = &activitycomm.Activity{
			Period: &activitycomm.Interval{
				From: from,
				To:   to,
			},
			Id:      act.ID,
			ActType: actType,
			Order:   order,
		}
	}

	return activities
}

func (s *activityServer) getOrders(ctx context.Context, acts storage.Activities) chan map[string]*ordercomm.Order {
	ordChan := make(chan map[string]*ordercomm.Order)
	go func() {
		defer close(ordChan)

		_ids := map[string]struct{}{}
		for _, act := range acts {
			_ids[act.OrderID] = struct{}{}
		}
		ids := []string{}
		for k := range _ids {
			if k != "" {
				ids = append(ids, k)
			}
		}

		ordersReply, err := s.orderCli.List(ctx,
			&ordercomm.ListOrdersRequest{FilterIds: ids})

		if err != nil {
			log.Printf("error retrieving the filtered list of orders: %v", err)
		} else {
			orders := make(map[string]*ordercomm.Order, len(ordersReply.Orders))
			for _, at := range ordersReply.Orders {
				orders[at.Id] = at
			}
			ordChan <- orders
		}
	}()

	return ordChan
}

func (s *activityServer) getActivities(ctx context.Context, acts storage.Activities) chan map[string]*activitytypecomm.ActivityType {
	atrChan := make(chan map[string]*activitytypecomm.ActivityType)
	go func() {
		defer close(atrChan)

		_ids := map[string]struct{}{}
		for _, act := range acts {
			_ids[act.ActivityTypeID] = struct{}{}
		}
		ids := []string{}
		for k := range _ids {
			if k != "" {
				ids = append(ids, k)
			}
		}

		activityTypesReply, err := s.actTypeCli.List(ctx,
			&activitytypecomm.ListActivityTypesRequest{FilterIds: ids})

		if err != nil {
			log.Printf("error retrieving the filtered list of activity types: %v", err)
		} else {
			actTypes := make(map[string]*activitytypecomm.ActivityType, len(activityTypesReply.ActivityTypes))
			for _, at := range activityTypesReply.ActivityTypes {
				actTypes[at.Id] = at
			}
			atrChan <- actTypes
		}
	}()

	return atrChan
}
