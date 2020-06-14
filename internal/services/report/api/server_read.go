package api

import (
	"context"
	"sort"

	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/reportgrpc"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *reportServer) Read(ctx context.Context, req *reportgrpc.ReadReportRequest) (*reportgrpc.ReadReportReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	report, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	activities := make([]*bookmastergrpc.Activity, len(report.Activities))
	for idx, act := range report.Activities {
		from, _ := ptypes.TimestampProto(act.Period.From)
		to, _ := ptypes.TimestampProto(act.Period.To)
		activities[idx] = &bookmastergrpc.Activity{
			Id: act.ID,
			Order: &ordergrpc.Order{
				Id:   act.Order.ID,
				Name: act.Order.Name,
			},
			ActType: &bookmastergrpc.ActivityType{
				Id:   act.Type.ID,
				Name: act.Type.Name,
				Code: int32(act.Type.Code),
			},
			Period: &bookmastergrpc.Interval{
				From: from,
				To:   to,
			},
		}
	}

	sort.Slice(activities, func(i, j int) bool {
		ifrom := activities[i].GetPeriod().GetFrom()
		jfrom := activities[j].GetPeriod().GetFrom()
		isec, jsec := ifrom.GetSeconds(), jfrom.GetSeconds()

		return isec < jsec ||
			(isec == jsec && ifrom.GetNanos() < jfrom.GetNanos())
	})

	from, _ := ptypes.TimestampProto(report.Period.From)
	to, _ := ptypes.TimestampProto(report.Period.To)
	return &reportgrpc.ReadReportReply{
		Report: &reportgrpc.Report{
			Name:       report.Name,
			Id:         report.ID,
			Activities: activities,
			Period: &reportgrpc.Interval{
				From: from,
				To:   to,
			},
		},
	}, nil
}
