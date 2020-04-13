package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type activityServer struct {
	repo        storage.Repository
	acttypehost string
}

// NewActivityServer returns the default implementation of ActivitySvcServer
func NewActivityServer(actTypeHost string) activitycomm.ActivitySvcServer {
	return &activityServer{
		repo:        storage.NewInMemoryStore(),
		acttypehost: actTypeHost,
	}
}

func (s *activityServer) Create(ctx context.Context, req *activitycomm.CreateActivityRequest) (*activitycomm.CreateActivityReply, error) {
	act := storage.Activity{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	id, err := s.repo.Create(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating activity: %v", err)
	}

	return &activitycomm.CreateActivityReply{
		Id: id.String(),
	}, nil
}

func (s *activityServer) Delete(ctx context.Context, req *activitycomm.DeleteActivityRequest) (*activitycomm.DeleteActivityReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Delete(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &activitycomm.DeleteActivityReply{
		Activity: &activitycomm.Activity{
			Code:        act.Code,
			Description: act.Description,
			Name:        act.Name,
			Id:          act.ID.String(),
		},
	}, nil
}

func (s *activityServer) Update(ctx context.Context, req *activitycomm.UpdateActivityRequest) (*activitycomm.UpdateActivityReply, error) {
	act := storage.Activity{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	uact, err := s.repo.Update(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating activity: %v", err)
	}

	return &activitycomm.UpdateActivityReply{
		Activity: &activitycomm.Activity{
			Code:        uact.Code,
			Description: uact.Description,
			Name:        uact.Name,
			Id:          uact.ID.String(),
		},
	}, nil
}

func (s *activityServer) List(ctx context.Context, req *activitycomm.ListActivitiesRequest) (*activitycomm.ListActivitiesReply, error) {
	acts, err := s.repo.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of activities: %v", err)
	}

	activities := []*activitycomm.Activity{}
	for _, v := range acts {
		activities = append(activities, &activitycomm.Activity{
			Code:        v.Code,
			Description: v.Description,
			Name:        v.Name,
			Id:          v.ID.String(),
		})
	}

	return &activitycomm.ListActivitiesReply{
		Activities: activities,
	}, nil
}
