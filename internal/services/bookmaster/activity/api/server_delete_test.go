package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/api"
	actteststore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage/teststore"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_DeleteHappyPath(t *testing.T) {
	// arrange
	activityID := uuid.New().String()
	store := &actteststore.ActivityTestRepo{
		DeleteResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityServer(store, nil, nil, 1*time.Second)
	ctx := context.Background()

	// act
	_, err := svr.Delete(ctx,
		&bookmastergrpc.DeleteActivityRequest{Id: activityID})

	// assert
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
}

func Test_DeleteInvalidID(t *testing.T) {
	// arrange
	store := &actteststore.ActivityTestRepo{}
	svr := api.NewActivityServer(store, nil, nil, 1*time.Second)

	// act
	_, err := svr.Delete(context.Background(),
		&bookmastergrpc.DeleteActivityRequest{Id: ""})

	// assert
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
