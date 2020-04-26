package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *reportServer) Create(ctx context.Context, req *reportcomm.CreateReportRequest) (*reportcomm.CreateReportReply, error) {
	period := req.GetPeriod()
	if period == nil || period.GetFrom() == nil || period.GetTo() == nil {
		return nil, status.Errorf(codes.InvalidArgument, `error creating report: provided period is null or invalid`)
	}

	rep, err := s.activityCli.ListInInterval(ctx,
		&activitycomm.ListInIntervalActivitiesRequest{
			Period: &activitycomm.Interval{
				From: period.GetFrom(),
				To:   period.GetTo(),
			},
		})
	if err != nil {
		return nil, err
	}

	acts := make(storage.Activities, len(rep.Activities))
	if len(rep.GetActivities()) > 0 {
		for idx, act := range rep.Activities {
			from, _ := ptypes.Timestamp(act.GetPeriod().GetFrom())
			to, _ := ptypes.Timestamp(act.GetPeriod().GetTo())

			acts[idx] = storage.Activity{
				Order: storage.Order{
					ID:   act.GetOrder().GetId(),
					Name: act.GetOrder().GetName(),
				},
				ID: act.GetId(),
				Period: storage.Interval{
					From: from,
					To:   to,
				},
				Type: storage.ActivityType{
					ID:   act.GetActType().GetId(),
					Code: int(act.GetActType().GetCode()),
					Name: act.GetActType().GetName(),
				},
			}
		}
	}

	report := storage.Report{
		Name:       req.GetName(),
		Activities: acts,
	}

	id, err := s.repo.Create(ctx, report)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating report: %v", err)
	}

	return &reportcomm.CreateReportReply{
		Id: id.String(),
	}, nil
}
