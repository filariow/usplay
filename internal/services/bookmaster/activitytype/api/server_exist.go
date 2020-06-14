package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityTypeServer) Exist(ctx context.Context, req *bookmastergrpc.ExistActivityTypeRequest) (*bookmastergrpc.ExistActivityTypeReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	exists, err := s.repo.Exist(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &bookmastergrpc.ExistActivityTypeReply{Exists: *exists}, nil
}
