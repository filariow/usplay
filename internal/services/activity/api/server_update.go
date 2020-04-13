package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) Update(ctx context.Context, req *activitycomm.UpdateActivityRequest) (*activitycomm.UpdateActivityReply, error) {
	// parse input
	idStr := req.GetId()
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error updating activity: id (%s) is invalid: %v", idStr, err)
	}

	actTypeIDStr := req.GetActTypeID()
	actTypeID, err := uuid.Parse(actTypeIDStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error updating activity: ActivityType ID (%s) is invalid: %v", actTypeID, err)
	}

	// build storage payload
	act := storage.Activity{
		ID:             id,
		Name:           req.GetName(),
		Code:           req.GetCode(),
		Description:    req.GetDescription(),
		ActivityTypeID: actTypeID,
	}

	// persist data in storage
	if err := s.repo.Update(ctx, act); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating activity: %v", err)
	}
	return &activitycomm.UpdateActivityReply{}, nil
}
