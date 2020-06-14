package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityTypeServer) Update(ctx context.Context, req *bookmastergrpc.UpdateActivityTypeRequest) (*bookmastergrpc.UpdateActivityTypeReply, error) {
	if _, err := uuid.Parse(req.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id provided (%v): %v", req.GetId(), err)
	}

	act := storage.ActivityType{
		ID:   req.GetId(),
		Name: req.GetName(),
		Code: req.GetCode(),
	}

	err := s.repo.Update(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating ActivityType: %v", err)
	}

	return &bookmastergrpc.UpdateActivityTypeReply{}, nil
}
