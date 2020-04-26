package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *reportServer) Read(ctx context.Context, req *reportcomm.ReadReportRequest) (*reportcomm.ReadReportReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	report, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	activities := make([]*activitycomm.Activity, len(report.Activities))
	for idx, act := range report.Activities {
		from, _ := ptypes.TimestampProto(act.Period.From)
		to, _ := ptypes.TimestampProto(act.Period.To)
		activities[idx] = &activitycomm.Activity{
			Id: act.ID,
			Order: &ordercomm.Order{
				Id:   act.Order.ID,
				Name: act.Order.Name,
			},
			ActType: &activitytypecomm.ActivityType{
				Id:   act.Type.ID,
				Name: act.Type.Name,
				Code: int32(act.Type.Code),
			},
			Period: &activitycomm.Interval{
				From: from,
				To:   to,
			},
		}
	}

	from, _ := ptypes.TimestampProto(report.Period.From)
	to, _ := ptypes.TimestampProto(report.Period.To)
	return &reportcomm.ReadReportReply{
		Report: &reportcomm.Report{
			Name:       report.Name,
			Id:         report.ID,
			Activities: activities,
			Period: &reportcomm.Interval{
				From: from,
				To:   to,
			},
		},
	}, nil
}
