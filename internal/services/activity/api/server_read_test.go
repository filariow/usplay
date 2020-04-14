package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
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
