package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityTypeServer) Read(ctx context.Context, req *activitytypecomm.ReadActivityTypeRequest) (*activitytypecomm.ReadActivityTypeReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &activitytypecomm.ReadActivityTypeReply{
		ActivityType: &activitytypecomm.ActivityType{
			Code:        act.Code,
			Name:        act.Name,
			Id:          act.ID,
			Description: act.Description,
		},
	}, nil
}
