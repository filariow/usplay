package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *orderServer) Read(ctx context.Context, req *ordergrpc.ReadOrderRequest) (*ordergrpc.ReadOrderReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &ordergrpc.ReadOrderReply{
		Order: &ordergrpc.Order{
			Code:        act.Code,
			Description: act.Description,
			Name:        act.Name,
			Id:          act.ID,
		},
	}, nil
}
