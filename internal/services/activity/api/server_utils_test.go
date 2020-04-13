package api_test

import (
	"context"
	"fmt"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Activity Test repository
type activityTestRepo struct {
	CreateResult struct {
		ID  uuid.UUID
		Err error
	}
	ReadResult struct {
		Activity storage.Activity
		Err      error
	}
	UpdateResult struct {
		Err error
	}
}

// Create
func (r *activityTestRepo) Create(context.Context, storage.Activity) (uuid.UUID, error) {
	return r.CreateResult.ID, r.CreateResult.Err
}

// Read
func (r *activityTestRepo) Read(ctx context.Context, id uuid.UUID) (storage.Activity, error) {
	activity := r.ReadResult.Activity
	activity.ID = id
	return activity, r.ReadResult.Err
}

// Update
func (r *activityTestRepo) Update(context.Context, storage.Activity) error {
	return r.UpdateResult.Err
}

// Delete
func (r *activityTestRepo) Delete(context.Context, uuid.UUID) error {
	return fmt.Errorf("Not implemented")
}

// List
func (r *activityTestRepo) List(context.Context) (storage.Activities, error) {
	return nil, fmt.Errorf("Not implemented")
}

// ActivityType Test Client
type actTestClient struct {
	WaitTime   time.Duration
	ReadResult struct {
		Err   error
		Reply activitytypecomm.ReadActivityTypeReply
	}
	ExistResult struct {
		Err   error
		Reply activitytypecomm.ExistActivityTypeReply
	}
}

func (c *actTestClient) Create(ctx context.Context, in *activitytypecomm.CreateActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.CreateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Create method is not implemented")
}

// Exists an activity
func (c *actTestClient) Exist(ctx context.Context, in *activitytypecomm.ExistActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.ExistActivityTypeReply, error) {
	return &c.ExistResult.Reply, c.ExistResult.Err
}

// Reads an ActivityType
func (c *actTestClient) Read(ctx context.Context, in *activitytypecomm.ReadActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.ReadActivityTypeReply, error) {
	time.Sleep(c.WaitTime)

	reply := c.ReadResult.Reply
	reply.ActivityType.Id = in.Id
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
	return nil, status.Errorf(codes.Unimplemented, "List method is not implemented")
}
