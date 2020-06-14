package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) Update(ctx context.Context, req *bookmastergrpc.UpdateActivityRequest) (*bookmastergrpc.UpdateActivityReply, error) {
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
		exists, err := s.acttyperepo.Exist(ctx, actTypeID)
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, "error retrieving ActivityType from repo: %v", err)
		}
		if exists == nil || !*exists {
			return nil, status.Errorf(codes.NotFound, "ActivityType with id %s do not exists", actTypeIDStr)
		}
	}
	{
		rpl, err := s.orderCli.Exist(ctx, &ordergrpc.ExistOrderRequest{Id: orderIDStr})
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, "error contacting order service: %v", err)
		}
		if !rpl.Exists {
			return nil, status.Errorf(codes.NotFound, "order with id %s do not exists", orderIDStr)
		}
	}

	period := req.GetPeriod()
	if period == nil || period.GetFrom() == nil || period.GetTo() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "error updating activity: period is invalid")
	}

	from, err := ptypes.Timestamp(period.GetFrom())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error updating activity: period's from is invalid: %v", err)
	}
	to, err := ptypes.Timestamp(period.GetTo())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error updating activity: period's to is invalid: %v", err)
	}

	// build storage payload
	act := storage.Activity{
		ID: id.String(),
		Period: storage.Interval{
			From: from,
			To:   to,
		},
		OrderID:        orderIDStr,
		ActivityTypeID: actTypeIDStr,
	}

	// persist data in storage
	if err := s.actrepo.Update(ctx, act); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating activity: %v", err)
	}
	return &bookmastergrpc.UpdateActivityReply{}, nil
}
