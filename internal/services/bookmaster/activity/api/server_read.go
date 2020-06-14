package api

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) Read(ctx context.Context, req *bookmastergrpc.ReadActivityRequest) (*bookmastergrpc.ReadActivityReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	return s.read(ctx, uid)
}

func (s *activityServer) read(ctx context.Context, uid uuid.UUID) (*bookmastergrpc.ReadActivityReply, error) {
	at, err := s.readFromRepo(ctx, uid)
	if err != nil {
		return nil, err
	}

	actOut, actErr := s.getActivityType(ctx, uid)
	ordOut, ordErr := s.getOrder(ctx, at.OrderID)

	var order *ordergrpc.Order
	var acttype *bookmastergrpc.ActivityType
	if err, open := <-actErr; open {
		log.Printf(
			"error contacting ActivityType service for details of entity with id %s: %v",
			at.ActivityTypeID, err)

		acttype = &bookmastergrpc.ActivityType{
			Id: at.ActivityTypeID,
		}
	} else {
		acttype = <-actOut
	}

	if err, open := <-ordErr; open {
		log.Printf(
			"error contacting Order service for details of entity with id %s: %v",
			at.OrderID, err)

		order = &ordergrpc.Order{Id: at.OrderID}
	} else {
		order = <-ordOut
	}

	from, _ := ptypes.TimestampProto(at.Period.From)
	to, _ := ptypes.TimestampProto(at.Period.To)
	return &bookmastergrpc.ReadActivityReply{
		Activity: &bookmastergrpc.Activity{
			Id:      at.ID,
			ActType: acttype,
			Order:   order,
			Period: &bookmastergrpc.Interval{
				From: from,
				To:   to,
			},
		},
	}, nil
}

func (s *activityServer) getOrder(ctx context.Context, uid string) (<-chan *ordergrpc.Order, <-chan error) {
	c := make(chan *ordergrpc.Order, 1)
	e := make(chan error, 1)

	go func() {
		resp, err := s.orderCli.Read(ctx, &ordergrpc.ReadOrderRequest{Id: uid})
		if err != nil {
			e <- err
		} else {
			c <- resp.GetOrder()
		}

		close(e)
		close(c)
	}()

	return c, e
}

func (s *activityServer) getActivityType(ctx context.Context, uid uuid.UUID) (<-chan *bookmastergrpc.ActivityType, chan error) {
	c := make(chan *bookmastergrpc.ActivityType, 1)
	e := make(chan error, 1)

	go func() {
		act, err := s.acttyperepo.Read(ctx, uid)
		if err != nil {
			e <- err
		}
		if act != nil {
			c <- &bookmastergrpc.ActivityType{
				Id:         act.ID,
				Name:       act.Name,
				Code:       act.Code,
				NeedsOrder: act.NeedsOrder,
			}
		}

		close(e)
		close(c)
	}()

	return c, e
}

func (s *activityServer) readFromRepo(ctx context.Context, uid uuid.UUID) (*storage.Activity, error) {
	act, err := s.actrepo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", uid.String())
	}
	return act, nil
}
