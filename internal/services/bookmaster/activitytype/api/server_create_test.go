package api_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/api"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage/teststore"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/google/uuid"
)

func Test_CreateHappyPath(t *testing.T) {
	// arrange
	id := uuid.New()
	store := &teststore.ActivityTypeTestRepo{
		CreateResult: struct {
			ID  *uuid.UUID
			Err error
		}{
			ID:  &id,
			Err: nil,
		},
	}
	svr := api.NewActivityTypeServer(store)
	ctx := context.Background()

	// act
	reply, err := svr.Create(ctx, &bookmastergrpc.CreateActivityTypeRequest{
		Code: int32(23),
		Name: "name",
	})

	// assert
	if err != nil {
		t.Fatalf("error invoking create: %v", err)
	}

	if expected, provided := store.CreateResult.ID.String(), reply.GetId(); expected != provided {
		t.Errorf("expected id %s, provided %s", expected, provided)
	}
}
