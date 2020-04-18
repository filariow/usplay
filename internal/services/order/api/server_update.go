package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *orderServer) Update(ctx context.Context, req *ordercomm.UpdateOrderRequest) (*ordercomm.UpdateOrderReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil || uid == uuid.Nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act := storage.Order{
		ID:          uid,
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	if err := s.repo.Update(ctx, act); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating order: %v", err)
	}

	return &ordercomm.UpdateOrderReply{}, nil
}
