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

func Test_DeleteHappyPath(t *testing.T) {
	// arrange
	activitytype := storage.ActivityType{
		ID:   uuid.New().String(),
		Code: 32,
		Name: "Name",
	}
	store := &teststore.ActivityTypeTestRepo{
		DeleteResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityTypeServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Delete(ctx,
		&bookmastergrpc.DeleteActivityTypeRequest{Id: activitytype.ID})

	// assert
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
}

func Test_DeleteEmptyID(t *testing.T) {
	// arrange
	store := &teststore.ActivityTypeTestRepo{}
	svr := api.NewActivityTypeServer(store)

	// act
	_, err := svr.Delete(context.Background(),
		&bookmastergrpc.DeleteActivityTypeRequest{Id: ""})
	if err == nil {
		t.Fatalf("expected error not returned: %v", err)
	}

	// assert
	if statusErr := status.Convert(err); statusErr == nil {
		t.Fatalf("error is not a status.Status error")
	} else {
		if providedCode := statusErr.Code(); providedCode != codes.InvalidArgument {
			t.Errorf("expected error code InvalidArgument, provided %v", providedCode)
		}
	}
}

func Test_DeleteInvalidID(t *testing.T) {
	// arrange
	store := &teststore.ActivityTypeTestRepo{}
	svr := api.NewActivityTypeServer(store)

	// act
	_, err := svr.Delete(context.Background(),
		&bookmastergrpc.DeleteActivityTypeRequest{Id: "19287asda"})
	if err == nil {
		t.Fatalf("expected error not returned: %v", err)
	}

	// assert
	if statusErr := status.Convert(err); statusErr == nil {
		t.Fatalf("error is not a status.Status error")
	} else {
		if providedCode := statusErr.Code(); providedCode != codes.InvalidArgument {
			t.Errorf("expected error code InvalidArgument, provided %v", providedCode)
		}
	}
}
