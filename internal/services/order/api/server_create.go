package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *orderServer) Create(ctx context.Context, req *ordercomm.CreateOrderRequest) (*ordercomm.CreateOrderReply, error) {
	act := storage.Order{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	id, err := s.repo.Create(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating order: %v", err)
	}

	return &ordercomm.CreateOrderReply{
		Id: id.String(),
	}, nil
}
