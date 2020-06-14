package api

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) List(ctx context.Context, req *bookmastergrpc.ListActivitiesRequest) (*bookmastergrpc.ListActivitiesReply, error) {
	ids := make([]uuid.UUID, len(req.FilterIds))
	for idx, i := range req.FilterIds {
		id, err := uuid.Parse(i)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "ID is not valid: %v", err)
		}
		ids[idx] = id
	}

	acts, err := s.actrepo.List(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of activities: %v", err)
	}

	if len(acts) == 0 {
		return &bookmastergrpc.ListActivitiesReply{Activities: nil}, nil
	}

	// collecting the needed ActivityType details
	atrChan, ordChan := s.getActivities(ctx, acts), s.getOrders(ctx, acts)
	orders, actTypes := <-ordChan, <-atrChan

	// building response with optional data
	activities := tobookmastergrpc(acts, orders, actTypes)

	return &bookmastergrpc.ListActivitiesReply{
		Activities: activities,
	}, nil
}

func tobookmastergrpc(acts storage.Activities, orders map[string]*ordergrpc.Order,
	actTypes map[string]*bookmastergrpc.ActivityType) []*bookmastergrpc.Activity {
	activities := make([]*bookmastergrpc.Activity, len(acts))
	for idx, act := range acts {
		// if info about the ActivityType was not retrieved, add the info you have
		actType, ok := actTypes[act.ActivityTypeID]
		if !ok {
			actType = &bookmastergrpc.ActivityType{Id: act.ActivityTypeID}
		}

		order, ok := orders[act.OrderID]
		if !ok {
			order = &ordergrpc.Order{Id: act.OrderID}
		}

		from, _ := ptypes.TimestampProto(act.Period.From)
		to, _ := ptypes.TimestampProto(act.Period.To)
		activities[idx] = &bookmastergrpc.Activity{
			Period: &bookmastergrpc.Interval{
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

func (s *activityServer) getOrders(ctx context.Context, acts storage.Activities) chan map[string]*ordergrpc.Order {
	ordChan := make(chan map[string]*ordergrpc.Order)
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
			&ordergrpc.ListOrdersRequest{FilterIds: ids})

		if err != nil {
			log.Printf("error retrieving the filtered list of orders: %v", err)
		} else {
			orders := make(map[string]*ordergrpc.Order, len(ordersReply.Orders))
			for _, at := range ordersReply.Orders {
				orders[at.Id] = at
			}
			ordChan <- orders
		}
	}()

	return ordChan
}

func (s *activityServer) getActivities(ctx context.Context, acts storage.Activities) chan map[string]*bookmastergrpc.ActivityType {
	atrChan := make(chan map[string]*bookmastergrpc.ActivityType)
	go func() {
		defer close(atrChan)

		_ids := map[string]struct{}{}
		for _, act := range acts {
			_ids[act.ActivityTypeID] = struct{}{}
		}
		ids := []uuid.UUID{}
		for k := range _ids {
			if id, err := uuid.Parse(k); err != nil {
				ids = append(ids, id)
			}
		}

		// TODO: retrieve from store activity types details
		acttypes, err := s.acttyperepo.List(ctx, ids)
		if err != nil {
			logrus.Errorf("error retrieving filtered list of activitytypes: %v", err)
			return
		}

		actTypes := make(map[string]*bookmastergrpc.ActivityType, len(acttypes))
		for _, at := range acttypes {
			actTypes[at.ID] = &bookmastergrpc.ActivityType{
				Id:         at.ID,
				Code:       at.Code,
				Name:       at.Name,
				NeedsOrder: at.NeedsOrder,
			}
		}
		atrChan <- actTypes
	}()

	return atrChan
}
