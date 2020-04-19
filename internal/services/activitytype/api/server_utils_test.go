package api_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ActivityType Test repository
type activityTypeTestRepo struct {
	CreateResult struct {
		ID  *uuid.UUID
		Err error
	}
	DeleteResult struct {
		Err error
	}
	ReadResult struct {
		ActivityType *storage.ActivityType
		Err          error
	}
	ExistResult struct {
		Result *bool
		Err    error
	}
	ListResult struct {
		Activities []storage.ActivityType
		Err        error
	}
	UpdateResult struct {
		Err error
	}
}

// Create
func (r *activityTypeTestRepo) Create(context.Context, storage.ActivityType) (*uuid.UUID, error) {
	return r.CreateResult.ID, r.CreateResult.Err
}

// Exist
func (r *activityTypeTestRepo) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	return r.ExistResult.Result, r.ExistResult.Err
}

// Read
func (r *activityTypeTestRepo) Read(ctx context.Context, id uuid.UUID) (*storage.ActivityType, error) {
	return r.ReadResult.ActivityType, r.ReadResult.Err
}

// Update
func (r *activityTypeTestRepo) Update(context.Context, storage.ActivityType) error {
	return r.UpdateResult.Err
}

// Delete
func (r *activityTypeTestRepo) Delete(context.Context, uuid.UUID) error {
	return r.DeleteResult.Err
}

// List
func (r *activityTypeTestRepo) List(context.Context, []uuid.UUID) (storage.ActivityTypes, error) {
	return r.ListResult.Activities, r.ListResult.Err
}

// ActivityType Test Client
type actTestClient struct {
	WaitTime   time.Duration
	ReadResult struct {
		Err   error
		Reply activitytypecomm.ReadActivityTypeReply
	}
	ListResult struct {
		Err   error
		Reply activitytypecomm.ListActivityTypesReply
	}
	ExistResult struct {
		Err   error
		Reply activitytypecomm.ExistActivityTypeReply
	}
}

func (c *actTestClient) Create(ctx context.Context, in *activitytypecomm.CreateActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.CreateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Create method is not implemented")
}

// Exists an activityType
func (c *actTestClient) Exist(ctx context.Context, in *activitytypecomm.ExistActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.ExistActivityTypeReply, error) {
	return &c.ExistResult.Reply, c.ExistResult.Err
}

// Reads an ActivityType
func (c *actTestClient) Read(ctx context.Context, in *activitytypecomm.ReadActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.ReadActivityTypeReply, error) {
	time.Sleep(c.WaitTime)

	reply := c.ReadResult.Reply
	if reply.GetActivityType() != nil {
		reply.ActivityType.Id = in.Id
	}
	return &reply, c.ReadResult.Err
}

// Delete an ActivityType
func (c *actTestClient) Delete(ctx context.Context, in *activitytypecomm.DeleteActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.DeleteActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Delete method is not implemented")
}

// Update an ActivityType
func (c *actTestClient) Update(ctx context.Context, in *activitytypecomm.UpdateActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.UpdateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Update method is not implemented")
}

// List an ActivityType
func (c *actTestClient) List(ctx context.Context, in *activitytypecomm.ListActivityTypesRequest, opts ...grpc.CallOption) (*activitytypecomm.ListActivityTypesReply, error) {
	return &c.ListResult.Reply, c.ListResult.Err
}

func assertEqActivityType(t *testing.T, expected, provided *activitytypecomm.ActivityType) {
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
