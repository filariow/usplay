package api

import (
	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
)

type orderServer struct {
	repo storage.Repository
}

// NewOrderServer returns the default implementation of OrderSvcServer
func NewOrderServer(repo storage.Repository) ordergrpc.OrderSvcServer {
	return &orderServer{repo: repo}
}
