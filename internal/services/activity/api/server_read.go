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

	act, err := s.getActivityType(ctx, at.ActivityTypeID)
	if err != nil {
		return nil, err
	}

	ord, err := s.getOrder(ctx, at.OrderID)
	if err != nil {
		return nil, err
	}

	return &activitycomm.ReadActivityReply{
		Activity: &activitycomm.Activity{
			Code:         at.Code,
			Description:  at.Description,
			Id:           at.ID.String(),
			Name:         at.Name,
			ActType:      act,
			CreationTime: creationTime,
			Order:        ord,
		},
	}, nil
}

func (s *activityServer) getOrder(ctx context.Context, uid uuid.UUID) (*ordercomm.Order, error) {
	resp, err := s.orderCli.Read(ctx, &ordercomm.ReadOrderRequest{Id: uid.String()})
	return resp.GetOrder(), err
}

func (s *activityServer) getActivityType(ctx context.Context, uid uuid.UUID) (*activitytypecomm.ActivityType, error) {
	resp, err := s.actTypeCli.Read(ctx, &activitytypecomm.ReadActivityTypeRequest{Id: uid.String()})
	return resp.GetActivityType(), err
}

func (s *activityServer) readFromRepo(ctx context.Context, uid uuid.UUID) (*storage.Activity, error) {
	act, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", uid.String())
	}
	return &act, nil
}
