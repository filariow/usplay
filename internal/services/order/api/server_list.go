package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *orderServer) List(ctx context.Context, req *ordercomm.ListOrdersRequest) (*ordercomm.ListOrdersReply, error) {
	acts, err := s.repo.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of orders: %v", err)
	}

	orders := []*ordercomm.Order{}
	for _, v := range acts {
		orders = append(orders, &ordercomm.Order{
			Code:        v.Code,
			Description: v.Description,
			Name:        v.Name,
			Id:          v.ID.String(),
		})
	}

	return &ordercomm.ListOrdersReply{
		Orders: orders,
	}, nil
}
