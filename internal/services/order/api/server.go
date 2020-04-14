package api

import (
	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
)

type orderServer struct {
	repo storage.Repository
}

// NewOrderServer returns the default implementation of OrderSvcServer
func NewOrderServer(repo storage.Repository) ordercomm.OrderSvcServer {
	return &orderServer{repo: repo}
}
