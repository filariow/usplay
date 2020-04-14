package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type reportServer struct {
	repo storage.Repository
}

// NewReportServer returns the default implementation of ReportSvcServer
func NewReportServer() reportcomm.ReportSvcServer {
	return &reportServer{
		repo: storage.NewInMemoryStore(),
	}
}

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
			Id:   report.ID.String(),
		},
	}, nil
}

func (s *reportServer) Delete(ctx context.Context, req *reportcomm.DeleteReportRequest) (*reportcomm.DeleteReportReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	report, err := s.repo.Delete(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &reportcomm.DeleteReportReply{
		Report: &reportcomm.Report{
			Name: report.Name,
			Id:   report.ID.String(),
		},
	}, nil
}

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
