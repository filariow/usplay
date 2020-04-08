package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/report/comm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type reportServer struct {
	repo storage.Repository
}

// NewReportServer returns the default implementation of ReportSvcServer
func NewReportServer() comm.ReportSvcServer {
	return &reportServer{
		repo: storage.NewInMemoryStore(),
	}
}

func (s *reportServer) Read(ctx context.Context, req *comm.ReadReportRequest) (*comm.ReadReportReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	report, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &comm.ReadReportReply{
		Report: &comm.Report{
			Name: report.Name,
			Id:   report.ID.String(),
		},
	}, nil
}

func (s *reportServer) Delete(ctx context.Context, req *comm.DeleteReportRequest) (*comm.DeleteReportReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	report, err := s.repo.Delete(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &comm.DeleteReportReply{
		Report: &comm.Report{
			Name: report.Name,
			Id:   report.ID.String(),
		},
	}, nil
}

func (s *reportServer) Update(ctx context.Context, req *comm.UpdateReportRequest) (*comm.UpdateReportReply, error) {
	report := storage.Report{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	uact, err := s.repo.Update(ctx, report)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating report: %v", err)
	}

	return &comm.UpdateReportReply{
		Report: &comm.Report{
			Name: uact.Name,
			Id:   uact.ID.String(),
		},
	}, nil
}

func (s *reportServer) List(ctx context.Context, req *comm.ListReportsRequest) (*comm.ListReportsReply, error) {
	sReports, err := s.repo.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of reports: %v", err)
	}

	reports := []*comm.Report{}
	for _, v := range sReports {
		reports = append(reports, &comm.Report{
			Name: v.Name,
			Id:   v.ID.String(),
		})
	}

	return &comm.ListReportsReply{
		Reports: reports,
	}, nil
}
