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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_ListHappyPath(t *testing.T) {
	// arrange
	activities := make([]storage.Activity, 2)
	activityTypes := make([]*activitytypecomm.ActivityType, 2)
	expectedActivities := make([]activitycomm.Activity, 2)

	for i := 0; i < 2; i++ {
		activityTypeID := uuid.New()
		activityType := activitytypecomm.ActivityType{
			Id:   activityTypeID.String(),
			Code: 1,
			Name: "Test ActivityType",
		}
		activityTypes[i] = &activityType

		activity := storage.Activity{
			ID:             uuid.New(),
			Code:           "Activity Code",
			Description:    "Activity Description",
			Name:           "Activity Name",
			ActivityTypeID: activityTypeID,
		}
		activities[i] = activity

		expectedActivity := activitycomm.Activity{
			Code:        activity.Code,
			Id:          activity.ID.String(),
			ActType:     &activityType,
			Description: activity.Description,
			Name:        activity.Name,
		}
		expectedActivities[i] = expectedActivity
	}

	store := &activityTestRepo{
		ListResult: struct {
			Activities []storage.Activity
			Err        error
		}{
			Activities: activities,
			Err:        nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ListResult: struct {
				Err   error
				Reply activitytypecomm.ListActivityTypesReply
			}{
				Err: nil,
				Reply: activitytypecomm.ListActivityTypesReply{
					ActivityTypes: activityTypes,
				},
			},
		},
		1*time.Second,
	)

	ctx := context.Background()

	// act
	reply, err := svr.List(ctx, &activitycomm.ListActivitiesRequest{})
	if err != nil {
		t.Fatalf("error invoking list: %v", err)
	}
	provActs := reply.GetActivities()

	// assert
	if provActs == nil {
		t.Fatalf("Returned activities is empty")
	}

	for idx, act := range provActs {
		assertEqActivity(t, &expectedActivities[idx], act)
		assertEqActivityType(t, activityTypes[idx], act.GetActType())
	}
}

func Test_ListHappyPathNoActivityTypesDetails(t *testing.T) {
	// arrange
	activities := make([]storage.Activity, 2)
	expectedActivities := make([]activitycomm.Activity, 2)

	for i := 0; i < 2; i++ {
		activityTypeID := uuid.New()

		activity := storage.Activity{
			ID:             uuid.New(),
			Code:           "Activity Code",
			Description:    "Activity Description",
			Name:           "Activity Name",
			ActivityTypeID: activityTypeID,
		}
		activities[i] = activity

		expectedActivity := activitycomm.Activity{
			Code: activity.Code,
			Id:   activity.ID.String(),
			ActType: &activitytypecomm.ActivityType{
				Id: activityTypeID.String(),
			},
			Description: activity.Description,
			Name:        activity.Name,
		}
		expectedActivities[i] = expectedActivity
	}

	store := &activityTestRepo{
		ListResult: struct {
			Activities []storage.Activity
			Err        error
		}{
			Activities: activities,
			Err:        nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ListResult: struct {
				Err   error
				Reply activitytypecomm.ListActivityTypesReply
			}{
				Err: status.Errorf(codes.DeadlineExceeded,
					"Could not retrieve details about Activity Types"),
				Reply: activitytypecomm.ListActivityTypesReply{},
			},
		},
		1*time.Second,
	)

	ctx := context.Background()

	// act
	reply, err := svr.List(ctx, &activitycomm.ListActivitiesRequest{})
	if err != nil {
		t.Fatalf("error invoking list: %v", err)
	}
	provActs := reply.GetActivities()

	// assert
	if provActs == nil {
		t.Fatalf("Returned activities is empty")
	}

	for idx, act := range provActs {
		assertEqActivity(t, &expectedActivities[idx], act)
		assertEqActivityType(t, expectedActivities[idx].ActType, act.ActType)
	}
}
