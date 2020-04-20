package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) Create(ctx context.Context, req *activitycomm.CreateActivityRequest) (*activitycomm.CreateActivityReply, error) {
	// parse inputs
	actTypeIDStr := req.GetActTypeID()
	_, err := uuid.Parse(actTypeIDStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error creating activity, ActivityType ID is invalid: %v", err)
	}

	// validate inputs
	rpl, err := s.actTypeCli.Exist(ctx, &activitytypecomm.ExistActivityTypeRequest{Id: actTypeIDStr})
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "error contacting ActivityType service: %v", err)
	}
	if !rpl.Exists {
		return nil, status.Errorf(codes.NotFound, "ActivityType with id %v do not exists", actTypeIDStr)
	}

	// create store data
	act := storage.Activity{
		Name:           req.GetName(),
		Code:           req.GetCode(),
		Description:    req.GetDescription(),
		ActivityTypeID: req.GetActTypeID(),
	}

	// persista data
	id, err := s.repo.Create(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating activity: %v", err)
	}

	// reply
	return &activitycomm.CreateActivityReply{
		Id: id.String(),
	}, nil
}
