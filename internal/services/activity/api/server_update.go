package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
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

	orderIDStr := req.GetOrderID()
	orderID, err := uuid.Parse(orderIDStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error updating activity: order ID (%s) is invalid: %v", orderID, err)
	}

	// validate inputs
	{
		rpl, err := s.actTypeCli.Exist(ctx, &activitytypecomm.ExistActivityTypeRequest{Id: actTypeIDStr})
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, "error contacting ActivityType service: %v", err)
		}
		if !rpl.Exists {
			return nil, status.Errorf(codes.NotFound, "ActivityType with id %s do not exists", actTypeIDStr)
		}
	}
	{
		rpl, err := s.orderCli.Exist(ctx, &ordercomm.ExistOrderRequest{Id: orderIDStr})
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, "error contacting order service: %v", err)
		}
		if !rpl.Exists {
			return nil, status.Errorf(codes.NotFound, "order with id %s do not exists", orderIDStr)
		}
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
