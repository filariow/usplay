package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) Delete(ctx context.Context, req *activitycomm.DeleteActivityRequest) (*activitycomm.DeleteActivityReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	if err := s.repo.Delete(ctx, uid); err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &activitycomm.DeleteActivityReply{}, nil
}
