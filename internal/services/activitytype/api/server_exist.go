package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityTypeServer) Exist(ctx context.Context, req *activitytypecomm.ExistActivityTypeRequest) (*activitytypecomm.ExistActivityTypeReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	exists, err := s.repo.Exist(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &activitytypecomm.ExistActivityTypeReply{Exists: *exists}, nil
}
