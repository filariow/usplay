package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *reportServer) Create(ctx context.Context, req *reportcomm.CreateReportRequest) (*reportcomm.CreateReportReply, error) {
	s.activityCli.List(ctx, &activitycomm.ListActivitiesRequest{})
	report := storage.Report{
		Name: req.GetName(),
	}

	id, err := s.repo.Create(ctx, report)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating report: %v", err)
	}

	return &reportcomm.CreateReportReply{
		Id: id.String(),
	}, nil
}
