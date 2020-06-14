package api

import (
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
)

type activityTypeServer struct {
	repo storage.Repository
}

// NewActivityTypeServer returns the default implementation of ActivityTypeSvcServer
func NewActivityTypeServer(repo storage.Repository) bookmastergrpc.ActivityTypeSvcServer {
	return &activityTypeServer{
		repo: repo,
	}
}
