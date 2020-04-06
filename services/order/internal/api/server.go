package api

import (
	"context"
	"usplay/us-order/gen"
	"usplay/us-order/pkg/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type orderServer struct {
	repo storage.Repository
}

// NewOrderServer returns the default implementation of OrderSvcServer
func NewOrderServer() gen.OrderSvcServer {
	return &orderServer{
		repo: storage.NewInMemoryStore(),
	}
}

func (s *orderServer) Create(ctx context.Context, req *gen.CreateOrderRequest) (*gen.CreateOrderReply, error) {
	act := storage.Order{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	id, err := s.repo.Create(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating order: %v", err)
	}

	return &gen.CreateOrderReply{
		Id: id.String(),
	}, nil
}

func (s *orderServer) Read(ctx context.Context, req *gen.ReadOrderRequest) (*gen.ReadOrderReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &gen.ReadOrderReply{
		Order: &gen.Order{
			Code:        act.Code,
			Description: act.Description,
			Name:        act.Name,
			Id:          act.ID.String(),
		},
	}, nil
}

func (s *orderServer) Delete(ctx context.Context, req *gen.DeleteOrderRequest) (*gen.DeleteOrderReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Delete(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &gen.DeleteOrderReply{
		Order: &gen.Order{
			Code:        act.Code,
			Description: act.Description,
			Name:        act.Name,
			Id:          act.ID.String(),
		},
	}, nil
}

func (s *orderServer) Update(ctx context.Context, req *gen.UpdateOrderRequest) (*gen.UpdateOrderReply, error) {
	act := storage.Order{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	uact, err := s.repo.Update(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating order: %v", err)
	}

	return &gen.UpdateOrderReply{
		Order: &gen.Order{
			Code:        uact.Code,
			Description: uact.Description,
			Name:        uact.Name,
			Id:          uact.ID.String(),
		},
	}, nil
}

func (s *orderServer) List(ctx context.Context, req *gen.ListActivitiesRequest) (*gen.ListActivitiesReply, error) {
	acts, err := s.repo.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of activities: %v", err)
	}

	activities := []*gen.Order{}
	for _, v := range acts {
		activities = append(activities, &gen.Order{
			Code:        v.Code,
			Description: v.Description,
			Name:        v.Name,
			Id:          v.ID.String(),
		})
	}

	return &gen.ListActivitiesReply{
		Activities: activities,
	}, nil
}
