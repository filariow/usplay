package api

import (
	"context"
	"usplay/us-activity/gen"
	"usplay/us-activity/pkg/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type activityServer struct {
	repo storage.Repository
}

// NewActivityServer returns the default implementation of ActivitySvcServer
func NewActivityServer() gen.ActivitySvcServer {
	return &activityServer{
		repo: storage.NewInMemoryStore(),
	}
}

func (s *activityServer) Create(ctx context.Context, req *gen.CreateActivityRequest) (*gen.CreateActivityReply, error) {
	act := storage.Activity{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	id, err := s.repo.Create(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating activity: %v", err)
	}

	return &gen.CreateActivityReply{
		Id: id.String(),
	}, nil
}

func (s *activityServer) Read(ctx context.Context, req *gen.ReadActivityRequest) (*gen.ReadActivityReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &gen.ReadActivityReply{
		Activity: &gen.Activity{
			Code:        act.Code,
			Description: act.Description,
			Name:        act.Name,
			Id:          act.ID.String(),
		},
	}, nil
}

func (s *activityServer) Delete(ctx context.Context, req *gen.DeleteActivityRequest) (*gen.DeleteActivityReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Delete(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &gen.DeleteActivityReply{
		Activity: &gen.Activity{
			Code:        act.Code,
			Description: act.Description,
			Name:        act.Name,
			Id:          act.ID.String(),
		},
	}, nil
}

func (s *activityServer) Update(ctx context.Context, req *gen.UpdateActivityRequest) (*gen.UpdateActivityReply, error) {
	act := storage.Activity{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	uact, err := s.repo.Update(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating activity: %v", err)
	}

	return &gen.UpdateActivityReply{
		Activity: &gen.Activity{
			Code:        uact.Code,
			Description: uact.Description,
			Name:        uact.Name,
			Id:          uact.ID.String(),
		},
	}, nil
}

func (s *activityServer) List(ctx context.Context, req *gen.ListActivitiesRequest) (*gen.ListActivitiesReply, error) {
	acts, err := s.repo.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of activities: %v", err)
	}

	activities := []*gen.Activity{}
	for _, v := range acts {
		activities = append(activities, &gen.Activity{
			Code:        v.Code,
			Description: v.Description,
			Name:        v.Name,
			Id:          v.ID.String(),
		})
	}

	return &gen.ListActivitiesReply{
		Activities: activities,
	}, nil
}
