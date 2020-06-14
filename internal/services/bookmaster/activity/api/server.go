package api

import (
	"time"

	actstore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	acttypestore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
)

type activityServer struct {
	waitTime    time.Duration
	actrepo     actstore.Repository
	acttyperepo acttypestore.Repository
	orderCli    ordergrpc.OrderSvcClient
}

// NewActivityServer returns the default implementation of ActivitySvcServer
func NewActivityServer(
	actrepo actstore.Repository,
	acttyperepo acttypestore.Repository,
	orderCli ordergrpc.OrderSvcClient,
	waitTime time.Duration) bookmastergrpc.ActivitySvcServer {
	return &activityServer{
		waitTime:    waitTime,
		actrepo:     actrepo,
		acttyperepo: acttyperepo,
		orderCli:    orderCli,
	}
}
