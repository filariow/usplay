package api

import (
	"github.com/FrancescoIlario/usplay/internal/services/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/reportgrpc"
)

type reportServer struct {
	activityCli bookmastergrpc.ActivitySvcClient
	repo        storage.Repository
}

// NewReportServer returns the default implementation of ReportSvcServer
func NewReportServer(client bookmastergrpc.ActivitySvcClient, repo storage.Repository) reportgrpc.ReportSvcServer {
	return &reportServer{
		activityCli: client,
		repo:        repo,
	}
}
