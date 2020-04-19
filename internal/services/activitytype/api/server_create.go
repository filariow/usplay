package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityTypeServer) Create(ctx context.Context, req *activitytypecomm.CreateActivityTypeRequest) (*activitytypecomm.CreateActivityTypeReply, error) {
	act := storage.ActivityType{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	id, err := s.repo.Create(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating ActivityType: %v", err)
	}

	return &activitytypecomm.CreateActivityTypeReply{
		Id: id.String(),
	}, nil
}
