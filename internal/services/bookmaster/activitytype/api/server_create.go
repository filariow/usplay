package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityTypeServer) Create(ctx context.Context, req *bookmastergrpc.CreateActivityTypeRequest) (*bookmastergrpc.CreateActivityTypeReply, error) {
	act := storage.ActivityType{
		Name: req.GetName(),
		Code: req.GetCode(),
	}

	id, err := s.repo.Create(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating ActivityType: %v", err)
	}

	return &bookmastergrpc.CreateActivityTypeReply{
		Id: id.String(),
	}, nil
}
