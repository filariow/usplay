package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type orderServer struct {
	repo storage.Repository
}

// NewOrderServer returns the default implementation of OrderSvcServer
func NewOrderServer() ordercomm.OrderSvcServer {
	return &orderServer{
		repo: storage.NewInMemoryStore(),
	}
}

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

func (s *orderServer) Read(ctx context.Context, req *ordercomm.ReadOrderRequest) (*ordercomm.ReadOrderReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &ordercomm.ReadOrderReply{
		Order: &ordercomm.Order{
			Code:        act.Code,
			Description: act.Description,
			Name:        act.Name,
			Id:          act.ID.String(),
		},
	}, nil
}

func (s *orderServer) Delete(ctx context.Context, req *ordercomm.DeleteOrderRequest) (*ordercomm.DeleteOrderReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	act, err := s.repo.Delete(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", id)
	}

	return &ordercomm.DeleteOrderReply{
		Order: &ordercomm.Order{
			Code:        act.Code,
			Description: act.Description,
			Name:        act.Name,
			Id:          act.ID.String(),
		},
	}, nil
}

func (s *orderServer) Update(ctx context.Context, req *ordercomm.UpdateOrderRequest) (*ordercomm.UpdateOrderReply, error) {
	act := storage.Order{
		Name:        req.GetName(),
		Code:        req.GetCode(),
		Description: req.GetDescription(),
	}

	uact, err := s.repo.Update(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating order: %v", err)
	}

	return &ordercomm.UpdateOrderReply{
		Order: &ordercomm.Order{
			Code:        uact.Code,
			Description: uact.Description,
			Name:        uact.Name,
			Id:          uact.ID.String(),
		},
	}, nil
}

func (s *orderServer) List(ctx context.Context, req *ordercomm.ListActivitiesRequest) (*ordercomm.ListActivitiesReply, error) {
	acts, err := s.repo.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of activities: %v", err)
	}

	activities := []*ordercomm.Order{}
	for _, v := range acts {
		activities = append(activities, &ordercomm.Order{
			Code:        v.Code,
			Description: v.Description,
			Name:        v.Name,
			Id:          v.ID.String(),
		})
	}

	return &ordercomm.ListActivitiesReply{
		Activities: activities,
	}, nil
}
