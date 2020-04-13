package api

import (
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
)

type activityServer struct {
	waitTime   time.Duration
	repo       storage.Repository
	actTypeCli activitytypecomm.ActivityTypeSvcClient
}

// NewActivityServer returns the default implementation of ActivitySvcServer
func NewActivityServer(
	repo storage.Repository,
	actTypeCli activitytypecomm.ActivityTypeSvcClient,
	waitTime time.Duration) activitycomm.ActivitySvcServer {
	return &activityServer{
		waitTime:   waitTime,
		repo:       repo,
		actTypeCli: actTypeCli,
	}
}
