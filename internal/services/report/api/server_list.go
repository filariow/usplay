package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *reportServer) List(ctx context.Context, req *reportcomm.ListReportsRequest) (*reportcomm.ListReportsReply, error) {
	sReports, err := s.repo.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of reports: %v", err)
	}

	reports := []*reportcomm.Report{}
	for _, v := range sReports {
		reports = append(reports, &reportcomm.Report{
			Name: v.Name,
			Id:   v.ID.String(),
		})
	}

	return &reportcomm.ListReportsReply{
		Reports: reports,
	}, nil
}
