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

func TestRead_HappyPath(t *testing.T) {
	// arrange
	actTypeId := uuid.New()

	activityTypeStorage := storage.ActivityType{
		ID:   actTypeId.String(),
		Code: 1,
		Name: "Test ActivityType",
	}

	req := bookmastergrpc.ActivityType{
		Id:   activityTypeStorage.ID,
		Code: activityTypeStorage.Code,
		Name: activityTypeStorage.Name,
	}

	store := &teststore.ActivityTypeTestRepo{
		ReadResult: struct {
			ActivityType *storage.ActivityType
			Err          error
		}{
			ActivityType: &activityTypeStorage,
			Err:          nil,
		},
	}
	svr := api.NewActivityTypeServer(store)

	ctx := context.Background()

	// act
	reply, err := svr.Read(ctx, &bookmastergrpc.ReadActivityTypeRequest{
		Id: req.Id,
	})
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}

	// assert
	assertEqActivityType(t, &req, reply.GetActivityType())
}
