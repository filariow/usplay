package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityTypeServer) Update(ctx context.Context, req *activitytypecomm.UpdateActivityTypeRequest) (*activitytypecomm.UpdateActivityTypeReply, error) {
	if _, err := uuid.Parse(req.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id provided (%v): %v", req.GetId(), err)
	}

	act := storage.ActivityType{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	err := s.repo.Update(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating ActivityType: %v", err)
	}

	return &activitytypecomm.UpdateActivityTypeReply{}, nil
}
