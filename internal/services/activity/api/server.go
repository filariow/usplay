package api

import (
	"context"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type activityServer struct {
	waitTime   time.Duration
	repo       storage.Repository
	actTypeCli activitytypecomm.ActivityTypeSvcClient
}

// NewActivityServer returns the default implementation of ActivitySvcServer
func NewActivityServer(
	repo storage.Repository,
	actTypeCli activitytypecomm.ActivityTypeSvcClient,
	waitTime time.Duration) activitycomm.ActivitySvcServer {
	return &activityServer{
		waitTime:   waitTime,
		repo:       repo,
		actTypeCli: actTypeCli,
	}
}

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
