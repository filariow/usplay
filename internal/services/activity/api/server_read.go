package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) Read(ctx context.Context, req *activitycomm.ReadActivityRequest) (*activitycomm.ReadActivityReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	return s.read(ctx, uid)
}

func (s *activityServer) read(ctx context.Context, uid uuid.UUID) (*activitycomm.ReadActivityReply, error) {
	at, err := s.readFromRepo(ctx, uid)
	if err != nil {
		return nil, err
	}
	creationTime, err := ptypes.TimestampProto(at.CreationTime)
	if err != nil {
		return nil, err
	}

	actOut, actErr := s.getActivityType(ctx, at.ActivityTypeID)
	ordOut, ordErr := s.getOrder(ctx, at.OrderID)
	if err, open := <-actErr; open {
		return nil, err
	}
	if err, open := <-ordErr; open {
		return nil, err
	}

	return &activitycomm.ReadActivityReply{
		Activity: &activitycomm.Activity{
			Code:         at.Code,
			Description:  at.Description,
			Id:           at.ID.String(),
			Name:         at.Name,
			ActType:      <-actOut,
			CreationTime: creationTime,
			Order:        <-ordOut,
		},
	}, nil
}

func (s *activityServer) getOrder(ctx context.Context, uid uuid.UUID) (<-chan *ordercomm.Order, <-chan error) {
	c := make(chan *ordercomm.Order, 1)
	e := make(chan error, 1)

	go func() {
		resp, err := s.orderCli.Read(ctx, &ordercomm.ReadOrderRequest{Id: uid.String()})
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

func (s *activityServer) getActivityType(ctx context.Context, uid uuid.UUID) (<-chan *activitytypecomm.ActivityType, chan error) {
	c := make(chan *activitytypecomm.ActivityType, 1)
	e := make(chan error, 1)

	go func() {
		resp, err := s.actTypeCli.Read(ctx, &activitytypecomm.ReadActivityTypeRequest{Id: uid.String()})
		if err != nil {
			e <- err
		} else {
			c <- resp.GetActivityType()
		}

		close(e)
		close(c)
	}()

	return c, e
}

func (s *activityServer) readFromRepo(ctx context.Context, uid uuid.UUID) (*storage.Activity, error) {
	act, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", uid.String())
	}
	return &act, nil
}
