package api_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
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

// Exists an activityType
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

func assertEqActivityType(t *testing.T, expected, provided *bookmastergrpc.ActivityType) {
	if proceed := assertNil(t, expected, provided); !proceed {
		return
	}

	if expected.Code != provided.Code {
		t.Errorf("expected activity type Code is %v while provided is %v", expected.Code, provided.Code)
	}
	if expected.Id != provided.Id {
		t.Errorf("expected activity type Id is %v while provided is %v", expected.Id, provided.Id)
	}
	if expected.Name != provided.Name {
		t.Errorf("expected activity type Name is %v while provided is %v", expected.Name, provided.Name)
	}
}

func isInterfaceNil(c interface{}) bool {
	return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())
}

func assertNil(t *testing.T, expected, provided interface{}) bool {
	expNil, provNil := isInterfaceNil(expected), isInterfaceNil(provided)

	if expNil && provNil {
		return false
	}

	if expNil || provNil {
		if !expNil {
			t.Errorf("provided value is nil while expected is not")
		} else {
			t.Errorf("expected value is nil while provided is not")
		}
		return false
	}
	return true
}
