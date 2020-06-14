package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/reportgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *reportServer) List(ctx context.Context, req *reportgrpc.ListReportsRequest) (*reportgrpc.ListReportsReply, error) {
	sReports, err := s.repo.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of reports: %v", err)
	}

	reports := []*reportgrpc.Report{}
	for _, v := range sReports {
		reports = append(reports, &reportgrpc.Report{
			Name: v.Name,
			Id:   v.ID,
		})
	}

	return &reportgrpc.ListReportsReply{
		Reports: reports,
	}, nil
}
