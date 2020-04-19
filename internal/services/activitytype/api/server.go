package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type activityTypeServer struct {
	repo storage.Repository
}

// NewActivityTypeServer returns the default implementation of ActivityTypeSvcServer
func NewActivityTypeServer(repo storage.Repository) activitytypecomm.ActivityTypeSvcServer {
	return &activityTypeServer{
		repo: repo,
	}
}

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

func (s *activityTypeServer) Read(ctx context.Context, req *activitytypecomm.ReadActivityTypeRequest) (*activitytypecomm.ReadActivityTypeReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &activitytypecomm.ReadActivityTypeReply{
		ActivityType: &activitytypecomm.ActivityType{
			Code:        act.Code,
			Name:        act.Name,
			Id:          act.ID,
			Description: act.Description,
		},
	}, nil
}

func (s *activityTypeServer) Delete(ctx context.Context, req *activitytypecomm.DeleteActivityTypeRequest) (*activitytypecomm.DeleteActivityTypeReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	if err := s.repo.Delete(ctx, uid); err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &activitytypecomm.DeleteActivityTypeReply{}, nil
}

func (s *activityTypeServer) Update(ctx context.Context, req *activitytypecomm.UpdateActivityTypeRequest) (*activitytypecomm.UpdateActivityTypeReply, error) {
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

func (s *activityTypeServer) List(ctx context.Context, req *activitytypecomm.ListActivityTypesRequest) (*activitytypecomm.ListActivityTypesReply, error) {
	ids := make([]uuid.UUID, len(req.FilterIds))
	for idx, i := range req.FilterIds {
		id, err := uuid.Parse(i)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "ID is not valid: %v", err)
		}
		ids[idx] = id
	}

	acts, err := s.repo.List(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of ActivityTypes: %v", err)
	}

	activityTypes := []*activitytypecomm.ActivityType{}
	for _, v := range acts {
		activityTypes = append(activityTypes, &activitytypecomm.ActivityType{
			Code: v.Code,
			Name: v.Name,
			Id:   v.ID,
		})
	}

	return &activitytypecomm.ListActivityTypesReply{
		ActivityTypes: activityTypes,
	}, nil
}

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
