package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *reportServer) Update(ctx context.Context, req *reportcomm.UpdateReportRequest) (*reportcomm.UpdateReportReply, error) {
	report := storage.Report{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	uact, err := s.repo.Update(ctx, report)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating report: %v", err)
	}

	return &reportcomm.UpdateReportReply{
		Report: &reportcomm.Report{
			Name: uact.Name,
			Id:   uact.ID.String(),
		},
	}, nil
}
