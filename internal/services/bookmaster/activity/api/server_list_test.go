package api_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	actteststore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage/teststore"
	acttypestorage "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	acttypeteststore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage/teststore"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
)

func Test_ListHappyPath(t *testing.T) {
	// arrange
	activities := make([]storage.Activity, 2)
	activityTypes := make([]acttypestorage.ActivityType, 2)
	activityTypesGrpc := make([]*bookmastergrpc.ActivityType, 2)
	orders := make([]*ordergrpc.Order, 2)
	expectedActivities := make([]bookmastergrpc.Activity, 2)

	for i := 0; i < 2; i++ {
		activityTypeID := uuid.New()
		activityType := acttypestorage.ActivityType{
			ID:         activityTypeID.String(),
			Code:       1,
			Name:       "Test ActivityType",
			NeedsOrder: i%2 == 0,
		}
		activityTypes[i] = activityType

		activityTypesGrpc[i] = &bookmastergrpc.ActivityType{
			Id:         activityType.ID,
			Code:       activityType.Code,
			Name:       activityType.Name,
			NeedsOrder: activityType.NeedsOrder,
		}

		activity := storage.Activity{
			ID: uuid.New().String(),
			Period: storage.Interval{
				From: time.Now(),
				To:   time.Now().Add(10 * 24 * time.Hour),
			},
			ActivityTypeID: activityTypeID.String(),
		}
		activities[i] = activity

		from, _ := ptypes.TimestampProto(activity.Period.From)
		to, _ := ptypes.TimestampProto(activity.Period.To)
		expectedActivities[i] = bookmastergrpc.Activity{
			Id:      activity.ID,
			ActType: activityTypesGrpc[i],
			Period: &bookmastergrpc.Interval{
				From: from,
				To:   to,
			},
		}

		order := ordergrpc.Order{
			Id:          uuid.New().String(),
			Code:        fmt.Sprintf("order code %v", i),
			Description: fmt.Sprintf("order description %v", i),
			Name:        fmt.Sprintf("order name %v", i),
		}
		orders[i] = &order
	}

	actstore := &actteststore.ActivityTestRepo{
		ListResult: struct {
			Activities []storage.Activity
			Err        error
		}{
			Activities: activities,
			Err:        nil,
		},
	}

	// TODO: complete here
	acttypestore := &acttypeteststore.ActivityTypeTestRepo{
		ListResult: struct {
			Activities []acttypestorage.ActivityType
			Err        error
		}{
			Activities: activityTypes,
			Err:        nil,
		},
	}

	svr := api.NewActivityServer(
		actstore,
		acttypestore,
		&orderTestClient{
			ListResult: struct {
				Err   error
				Reply ordergrpc.ListOrdersReply
			}{
				Err:   nil,
				Reply: ordergrpc.ListOrdersReply{Orders: orders},
			},
		},
		1*time.Second,
	)

	ctx := context.Background()

	// act
	reply, err := svr.List(ctx, &bookmastergrpc.ListActivitiesRequest{})
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
		assertEqActivityType(t, activityTypesGrpc[idx], act.GetActType())
	}
}

func Test_ListHappyPathNoActivityTypesDetails(t *testing.T) {
	// arrange
	activities := make([]storage.Activity, 2)
	expectedActivities := make([]bookmastergrpc.Activity, 2)
	orders := make([]*ordergrpc.Order, 2)

	for i := 0; i < 2; i++ {
		activityTypeID := uuid.New()

		activity := storage.Activity{
			ID: uuid.New().String(),
			Period: storage.Interval{
				From: time.Now(),
				To:   time.Now().Add(10 * 24 * time.Hour),
			},
			ActivityTypeID: activityTypeID.String(),
		}
		activities[i] = activity

		from, _ := ptypes.TimestampProto(activity.Period.From)
		to, _ := ptypes.TimestampProto(activity.Period.To)
		expectedActivities[i] = bookmastergrpc.Activity{
			Id: activity.ID,
			ActType: &bookmastergrpc.ActivityType{
				Id: activityTypeID.String(),
			},
			Period: &bookmastergrpc.Interval{
				From: from,
				To:   to,
			},
		}

		order := ordergrpc.Order{
			Id:          uuid.New().String(),
			Code:        fmt.Sprintf("order code %v", i),
			Description: fmt.Sprintf("order description %v", i),
			Name:        fmt.Sprintf("order name %v", i),
		}
		orders[i] = &order
	}

	actstore := &actteststore.ActivityTestRepo{
		ListResult: struct {
			Activities []storage.Activity
			Err        error
		}{
			Activities: activities,
			Err:        nil,
		},
	}
	// TODO: complete here
	acttypestore := &acttypeteststore.ActivityTypeTestRepo{}
	svr := api.NewActivityServer(
		actstore,
		acttypestore,
		&orderTestClient{
			ListResult: struct {
				Err   error
				Reply ordergrpc.ListOrdersReply
			}{
				Err:   nil,
				Reply: ordergrpc.ListOrdersReply{Orders: orders},
			},
		},
		1*time.Second,
	)

	ctx := context.Background()

	// act
	reply, err := svr.List(ctx, &bookmastergrpc.ListActivitiesRequest{})
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
