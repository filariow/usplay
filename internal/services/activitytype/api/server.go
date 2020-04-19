package api

import (
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
)

type activityTypeServer struct {
	repo storage.Repository
}

// NewActivityTypeServer returns the default implementation of ActivityTypeSvcServer
func NewActivityTypeServer(repo storage.Repository) activitytypecomm.ActivityTypeSvcServer {
	return &activityTypeServer{
		repo: repo,
	}
}
