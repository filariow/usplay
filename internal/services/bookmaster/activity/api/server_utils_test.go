package api_test

import (
	"context"
	"time"

	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ActivityType Test Client
type actTestClient struct {
	WaitTime   time.Duration
	ReadResult struct {
		Err   error
		Reply bookmastergrpc.ReadActivityTypeReply
	}
	ListResult struct {
		Err   error
		Reply bookmastergrpc.ListActivityTypesReply
	}
	ExistResult struct {
		Err   error
		Reply bookmastergrpc.ExistActivityTypeReply
	}
}

func (c *actTestClient) Create(ctx context.Context, in *bookmastergrpc.CreateActivityTypeRequest, opts ...grpc.CallOption) (*bookmastergrpc.CreateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Create method is not implemented")
}

// Exists an activity
func (c *actTestClient) Exist(ctx context.Context, in *bookmastergrpc.ExistActivityTypeRequest, opts ...grpc.CallOption) (*bookmastergrpc.ExistActivityTypeReply, error) {
	return &c.ExistResult.Reply, c.ExistResult.Err
}

// Reads an ActivityType
func (c *actTestClient) Read(ctx context.Context, in *bookmastergrpc.ReadActivityTypeRequest, opts ...grpc.CallOption) (*bookmastergrpc.ReadActivityTypeReply, error) {
	time.Sleep(c.WaitTime)

	reply := &c.ReadResult.Reply
	if reply.GetActivityType() != nil {
		reply.ActivityType.Id = in.Id
	}
	return reply, c.ReadResult.Err
}

// Delete an ActivityType
func (c *actTestClient) Delete(ctx context.Context, in *bookmastergrpc.DeleteActivityTypeRequest, opts ...grpc.CallOption) (*bookmastergrpc.DeleteActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Delete method is not implemented")
}

// Update an ActivityType
func (c *actTestClient) Update(ctx context.Context, in *bookmastergrpc.UpdateActivityTypeRequest, opts ...grpc.CallOption) (*bookmastergrpc.UpdateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Update method is not implemented")
}

// List an ActivityType
func (c *actTestClient) List(ctx context.Context, in *bookmastergrpc.ListActivityTypesRequest, opts ...grpc.CallOption) (*bookmastergrpc.ListActivityTypesReply, error) {
	return &c.ListResult.Reply, c.ListResult.Err
}

// Order Test Client
type orderTestClient struct {
	WaitTime   time.Duration
	ReadResult struct {
		Err   error
		Reply ordergrpc.ReadOrderReply
	}
	ListResult struct {
		Err   error
		Reply ordergrpc.ListOrdersReply
	}
	ExistResult struct {
		Err   error
		Reply ordergrpc.ExistOrderReply
	}
}

func (c *orderTestClient) Create(ctx context.Context, in *ordergrpc.CreateOrderRequest, opts ...grpc.CallOption) (*ordergrpc.CreateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Create method is not implemented")
}

// Exists an activity
func (c *orderTestClient) Exist(ctx context.Context, in *ordergrpc.ExistOrderRequest, opts ...grpc.CallOption) (*ordergrpc.ExistOrderReply, error) {
	return &c.ExistResult.Reply, c.ExistResult.Err
}

// Reads an Order
func (c *orderTestClient) Read(ctx context.Context, in *ordergrpc.ReadOrderRequest, opts ...grpc.CallOption) (*ordergrpc.ReadOrderReply, error) {
	time.Sleep(c.WaitTime)

	reply := &c.ReadResult.Reply
	if reply.GetOrder() != nil {
		reply.Order.Id = in.Id
	}
	return reply, c.ReadResult.Err
}

// Delete an Order
func (c *orderTestClient) Delete(ctx context.Context, in *ordergrpc.DeleteOrderRequest, opts ...grpc.CallOption) (*ordergrpc.DeleteOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Delete method is not implemented")
}

// Update an Order
func (c *orderTestClient) Update(ctx context.Context, in *ordergrpc.UpdateOrderRequest, opts ...grpc.CallOption) (*ordergrpc.UpdateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Update method is not implemented")
}

// List an Order
func (c *orderTestClient) List(ctx context.Context, in *ordergrpc.ListOrdersRequest, opts ...grpc.CallOption) (*ordergrpc.ListOrdersReply, error) {
	return &c.ListResult.Reply, c.ListResult.Err
}
