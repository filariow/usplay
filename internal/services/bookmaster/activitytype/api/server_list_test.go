package api_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/api"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage/teststore"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/google/uuid"
)

func Test_ListHappyPath(t *testing.T) {
	// arrange
	activityTypesComm := make([]bookmastergrpc.ActivityType, 2)
	activityTypesStorage := make([]storage.ActivityType, 2)

	for i := 0; i < 2; i++ {
		activityTypeID := uuid.New()
		activityTypesComm[i] = bookmastergrpc.ActivityType{
			Id:   activityTypeID.String(),
			Code: 1,
			Name: "Test ActivityType",
		}
		activityType := &activityTypesComm[i]

		activityTypesStorage[i] = storage.ActivityType{
			ID:   activityTypeID.String(),
			Code: int32(activityType.Code),
			Name: activityType.Name,
		}
	}

	store := &teststore.ActivityTypeTestRepo{
		ListResult: struct {
			Activities []storage.ActivityType
			Err        error
		}{
			Activities: activityTypesStorage,
			Err:        nil,
		},
	}
	svr := api.NewActivityTypeServer(store)
	ctx := context.Background()

	// act
	reply, err := svr.List(ctx, &bookmastergrpc.ListActivityTypesRequest{})
	if err != nil {
		t.Fatalf("error invoking list: %v", err)
	}
	provActs := reply.GetActivityTypes()

	// assert
	if provActs == nil {
		t.Fatalf("Returned activities is empty")
	}

	for idx, act := range provActs {
		assertEqActivityType(t, &activityTypesComm[idx], act)
	}
}
