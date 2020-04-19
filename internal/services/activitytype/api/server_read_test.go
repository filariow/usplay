package api_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/api"
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
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

	activityTypeComm := activitytypecomm.ActivityType{
		Id:          activityTypeStorage.ID,
		Code:        activityTypeStorage.Code,
		Name:        activityTypeStorage.Name,
		Description: activityTypeStorage.Description,
	}

	store := &activityTypeTestRepo{
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
	reply, err := svr.Read(ctx, &activitytypecomm.ReadActivityTypeRequest{
		Id: activityTypeComm.Id,
	})
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}

	// assert
	assertEqActivityType(t, &activityTypeComm, reply.GetActivityType())
}
