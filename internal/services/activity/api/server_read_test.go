package api_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestReadHappyPath(t *testing.T) {
	// arrange
	activity := storage.Activity{
		ID:          uuid.New(),
		Code:        "Activity Code",
		Description: "Activity Description",
		Name:        "Activity Name",
	}
	activityType := activitytypecomm.ActivityType{
		Id:   uuid.New().String(),
		Code: 1,
		Name: "Test ActivityType",
	}
	expectedActivity := activitycomm.Activity{
		Code:        activity.Code,
		Id:          activity.ID.String(),
		ActType:     &activityType,
		Description: activity.Description,
		Name:        activity.Name,
	}

	store := &activityTestRepo{
		ReadResult: struct {
			Activity storage.Activity
			Err      error
		}{
			Activity: activity,
			Err:      nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ReadResult: struct {
				Err   error
				Reply activitytypecomm.ReadActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ReadActivityTypeReply{
					ActivityType: &activityType,
				},
			},
		},
		1*time.Second,
	)

	ctx := context.Background()

	// act
	reply, err := svr.Read(ctx, &activitycomm.ReadActivityRequest{
		Id: activity.ID.String(),
	})
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
	act := reply.GetActivity()

	// assert
	if act == nil {
		t.Fatalf("Returned activity is empty")
	}

	assertEqActivity(t, &expectedActivity, act)
	assertEqActivityType(t, &activityType, act.GetActType())
}

func Test_ReadTimeout(t *testing.T) {
	// arrange
	actID := uuid.New()
	store := &activityTestRepo{
		ReadResult: struct {
			Activity storage.Activity
			Err      error
		}{
			Activity: storage.Activity{
				ID:          actID,
				Code:        "Activity Code",
				Description: "Activity Description",
				Name:        "Activity Name",
			},
			Err: nil,
		},
	}
	svr := api.NewActivityServer(store,
		&actTestClient{
			WaitTime: 200 * time.Millisecond,
			ReadResult: struct {
				Err   error
				Reply activitytypecomm.ReadActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ReadActivityTypeReply{
					ActivityType: &activitytypecomm.ActivityType{
						Id:   uuid.New().String(),
						Code: 1,
						Name: "Test ActivityType",
					},
				},
			},
		},
		100*time.Millisecond,
	)

	ctx := context.Background()
	// act
	if _, err := svr.Read(ctx, &activitycomm.ReadActivityRequest{Id: actID.String()}); err == nil {
		t.Fatalf("no error even with client in timeout")
	}
}

func assertNil(t *testing.T, expected, provided interface{}) bool {
	if expected == nil && provided == nil {
		return false
	}

	if expected == nil || provided == nil {
		if expected != nil {
			t.Errorf("provided value is nil while expected is not")
		} else {
			t.Errorf("expected value is nil while provided is not")
		}
		return false
	}
	return true
}

func assertEqActivity(t *testing.T, expected, provided *activitycomm.Activity) {
	if proceed := assertNil(t, expected, provided); !proceed {
		return
	}

	if expected.Code != provided.Code {
		t.Errorf("expected Code is %v while provided is %v", expected.Code, provided.Code)
	}
	if expected.Description != provided.Description {
		t.Errorf("expected Description is %v while provided is %v", expected.Description, provided.Description)
	}
	if expected.Id != provided.Id {
		t.Errorf("expected Id is %v while provided is %v", expected.Id, provided.Id)
	}
	if expected.Name != provided.Name {
		t.Errorf("expected Name is %v while provided is %v", expected.Name, provided.Name)
	}
}

func assertEqActivityType(t *testing.T, expected, provided *activitytypecomm.ActivityType) {
	if proceed := assertNil(t, expected, provided); !proceed {
		return
	}

	if expected.Code != provided.Code {
		t.Errorf("expected Code is %v while provided is %v", expected.Code, provided.Code)
	}
	if expected.Id != provided.Id {
		t.Errorf("expected Id is %v while provided is %v", expected.Id, provided.Id)
	}
	if expected.Name != provided.Name {
		t.Errorf("expected Name is %v while provided is %v", expected.Name, provided.Name)
	}
}

// Activity Test repository
type activityTestRepo struct {
	ReadResult struct {
		Activity storage.Activity
		Err      error
	}
}

// Create
func (r *activityTestRepo) Create(context.Context, storage.Activity) (uuid.UUID, error) {
	return uuid.Nil, fmt.Errorf("Not implemented")
}

// Read
func (r *activityTestRepo) Read(ctx context.Context, id uuid.UUID) (storage.Activity, error) {
	activity := r.ReadResult.Activity
	activity.ID = id
	return activity, r.ReadResult.Err
}

// Update
func (r *activityTestRepo) Update(context.Context, storage.Activity) (storage.Activity, error) {
	return storage.Activity{}, fmt.Errorf("Not implemented")
}

// Delete
func (r *activityTestRepo) Delete(context.Context, uuid.UUID) (storage.Activity, error) {
	return storage.Activity{}, fmt.Errorf("Not implemented")
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
}

func (c *actTestClient) Create(ctx context.Context, in *activitytypecomm.CreateActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.CreateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Create method is not implemented")
}

// Exists an activity
func (c *actTestClient) Exist(ctx context.Context, in *activitytypecomm.ExistActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.ExistActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Exist method is not implemented")
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
