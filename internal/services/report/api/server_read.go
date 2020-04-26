package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
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

	return &reportcomm.ReadReportReply{
		Report: &reportcomm.Report{
			Name: report.Name,
			Id:   report.ID,
		},
	}, nil
}
