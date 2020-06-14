package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *orderServer) List(ctx context.Context, req *ordergrpc.ListOrdersRequest) (*ordergrpc.ListOrdersReply, error) {
	ids := make([]uuid.UUID, len(req.FilterIds))
	for idx, i := range req.FilterIds {
		id, err := uuid.Parse(i)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "ID is not valid: %v", err)
		}
		ids[idx] = id
	}

	acts, err := s.repo.List(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of orders: %v", err)
	}

	orders := []*ordergrpc.Order{}
	for _, v := range acts {
		orders = append(orders, &ordergrpc.Order{
			Code:        v.Code,
			Description: v.Description,
			Name:        v.Name,
			Id:          v.ID,
		})
	}

	return &ordergrpc.ListOrdersReply{
		Orders: orders,
	}, nil
}
