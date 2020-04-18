package api

import (
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
)

type activityServer struct {
	waitTime   time.Duration
	repo       storage.Repository
	actTypeCli activitytypecomm.ActivityTypeSvcClient
	orderCli   ordercomm.OrderSvcClient
}

// NewActivityServer returns the default implementation of ActivitySvcServer
func NewActivityServer(
	repo storage.Repository,
	actTypeCli activitytypecomm.ActivityTypeSvcClient,
	orderCli ordercomm.OrderSvcClient,
	waitTime time.Duration) activitycomm.ActivitySvcServer {
	return &activityServer{
		waitTime:   waitTime,
		repo:       repo,
		actTypeCli: actTypeCli,
		orderCli:   orderCli,
	}
}
