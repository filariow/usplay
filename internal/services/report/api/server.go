package api

import (
	"github.com/FrancescoIlario/usplay/internal/services/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
)

type reportServer struct {
	activityCli activitycomm.ActivitySvcClient
	repo        storage.Repository
}

// NewReportServer returns the default implementation of ReportSvcServer
func NewReportServer(client activitycomm.ActivitySvcClient, repo storage.Repository) reportcomm.ReportSvcServer {
	return &reportServer{
		activityCli: client,
		repo:        repo,
	}
}
