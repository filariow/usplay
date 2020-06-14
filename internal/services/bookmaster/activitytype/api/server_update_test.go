package api_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/api"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage/teststore"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_UpdateHappyPath(t *testing.T) {
	// arrange
	id := uuid.New()
	activity := storage.ActivityType{
		ID:         id.String(),
		Code:       int32(1),
		NeedsOrder: false,
		Name:       "ActivityType Name",
	}
	store := &teststore.ActivityTypeTestRepo{
		UpdateResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityTypeServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Update(ctx, &bookmastergrpc.UpdateActivityTypeRequest{
		Id:   activity.ID,
		Code: activity.Code,
		Name: activity.Name,
	})

	// assert
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
}

func TestUpdate_InvalidActivityID(t *testing.T) {
	// arrange
	activity := storage.ActivityType{
		Code:       int32(2),
		NeedsOrder: false,
		Name:       "ActivityType Name",
	}
	store := &teststore.ActivityTypeTestRepo{
		UpdateResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityTypeServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Update(ctx, &bookmastergrpc.UpdateActivityTypeRequest{
		Id:         "",
		Code:       activity.Code,
		Name:       activity.Name,
		NeedsOrder: true,
	})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking update with no id is not provided")
	}

	statusErr := status.Convert(err)
	if statusErr == nil {
		t.Fatalf("provided error is not a status.Status error: %v", err)
	}

	if statusErr.Code() != codes.InvalidArgument {
		t.Errorf("provided error do not present the InvalidArgument code as expected, but instead presents %s", statusErr.Code().String())
	}
}
