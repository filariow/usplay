package api

import (
	"github.com/FrancescoIlario/usplay/internal/services/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
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
