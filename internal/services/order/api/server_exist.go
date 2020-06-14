package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *orderServer) Exist(ctx context.Context, req *ordergrpc.ExistOrderRequest) (*ordergrpc.ExistOrderReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id provided (%s): %v", id, err)
	}

	exists, err := s.repo.Exist(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &ordergrpc.ExistOrderReply{Exists: *exists}, nil
}
