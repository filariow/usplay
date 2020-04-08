package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/report/comm"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *reportServer) Generate(ctx context.Context, req *comm.GenerateReportRequest) (*comm.GenerateReportReply, error) {
	report := storage.Report{
		Name: req.GetName(),
	}

	id, err := s.repo.Create(ctx, report)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating report: %v", err)
	}

	return &comm.GenerateReportReply{
		Id: id.String(),
	}, nil
}
