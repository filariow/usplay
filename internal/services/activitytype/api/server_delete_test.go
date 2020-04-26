package api_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/api"
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
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
	store := &activityTypeTestRepo{
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
		&activitytypecomm.DeleteActivityTypeRequest{Id: activitytype.ID})

	// assert
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
}

func Test_DeleteEmptyID(t *testing.T) {
	// arrange
	store := &activityTypeTestRepo{}
	svr := api.NewActivityTypeServer(store)

	// act
	_, err := svr.Delete(context.Background(),
		&activitytypecomm.DeleteActivityTypeRequest{Id: ""})
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
	store := &activityTypeTestRepo{}
	svr := api.NewActivityTypeServer(store)

	// act
	_, err := svr.Delete(context.Background(),
		&activitytypecomm.DeleteActivityTypeRequest{Id: "19287asda"})
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
