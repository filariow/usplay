package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_DeleteHappyPath(t *testing.T) {
	// arrange
	activity := storage.Activity{
		ID:          uuid.New(),
		Code:        "Activity Code",
		Description: "Activity Description",
		Name:        "Activity Name",
	}
	store := &activityTestRepo{
		DeleteResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityServer(store, nil, 1*time.Second)
	ctx := context.Background()

	// act
	_, err := svr.Delete(ctx,
		&activitycomm.DeleteActivityRequest{Id: activity.ID.String()})

	// assert
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
}

func Test_DeleteInvalidID(t *testing.T) {
	// arrange
	store := &activityTestRepo{}
	svr := api.NewActivityServer(store, nil, 1*time.Second)

	// act

	// assert
	_, err := svr.Delete(context.Background(),
		&activitycomm.DeleteActivityRequest{Id: ""})
	if err == nil {
		t.Fatalf("expected error not returned: %v", err)
	}

	if statusErr := status.Convert(err); statusErr == nil {
		t.Fatalf("error is not a status.Status error")
	} else {
		if providedCode := statusErr.Code(); providedCode != codes.InvalidArgument {
			t.Errorf("expected error code InvalidArgument, provided %v", providedCode)
		}
	}
}
