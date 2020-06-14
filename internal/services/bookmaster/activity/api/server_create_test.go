package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	actteststore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage/teststore"
	acttypeteststore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage/teststore"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_CreateHappyPath(t *testing.T) {
	// arrange
	activityID := uuid.New()
	activity := storage.Activity{
		ActivityTypeID: uuid.New().String(),
		OrderID:        uuid.New().String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
	}
	actstore := &actteststore.ActivityTestRepo{
		CreateResult: struct {
			ID  *uuid.UUID
			Err error
		}{
			ID:  &activityID,
			Err: nil,
		},
	}
	res := true
	acttypestore := &acttypeteststore.ActivityTypeTestRepo{
		ExistResult: struct {
			Result *bool
			Err    error
		}{
			Result: &res, Err: nil,
		},
	}
	svr := api.NewActivityServer(
		actstore,
		acttypestore,
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordergrpc.ExistOrderReply
			}{
				Err: nil,
				Reply: ordergrpc.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()
	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.From)
	req := &bookmastergrpc.CreateActivityRequest{
		ActTypeID: activity.ActivityTypeID,
		OrderID:   activity.OrderID,
		Period: &bookmastergrpc.Interval{
			From: from,
			To:   to,
		},
	}

	// act
	reply, err := svr.Create(ctx, req)

	// assert
	if err != nil {
		t.Fatalf("error invoking create: %v", err)
	}

	if expected, provided := actstore.CreateResult.ID.String(), reply.GetId(); expected != provided {
		t.Errorf("expected id %s, provided %s", expected, provided)
	}
}

func Test_CreateInvalidActivityTypeID(t *testing.T) {
	// arrange
	activityID, orderID := uuid.New(), uuid.New()
	activity := storage.Activity{
		ID:      activityID.String(),
		OrderID: orderID.String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
	}
	actstore := &actteststore.ActivityTestRepo{
		CreateResult: struct {
			ID  *uuid.UUID
			Err error
		}{
			ID:  &activityID,
			Err: nil,
		},
	}
	// TODO: complete here
	acttypestore := &acttypeteststore.ActivityTypeTestRepo{}
	svr := api.NewActivityServer(
		actstore,
		acttypestore,
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordergrpc.ExistOrderReply
			}{
				Err: nil,
				Reply: ordergrpc.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()
	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.From)

	// act
	_, err := svr.Create(ctx, &bookmastergrpc.CreateActivityRequest{
		ActTypeID: "",
		Period: &bookmastergrpc.Interval{
			From: from,
			To:   to,
		},
	})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking create with no id is not provided")
	}

	statusErr := status.Convert(err)
	if statusErr == nil {
		t.Fatalf("provided error is not a status.Status error: %v", err)
	}

	if statusErr.Code() != codes.InvalidArgument {
		t.Errorf("provided error do not present the InvalidArgument code as expected, but instead presents %s", statusErr.Code().String())
	}
}

func Test_CreateNotExistingActivityTypeID(t *testing.T) {
	// arrange
	activityID := uuid.New()
	activity := storage.Activity{
		ActivityTypeID: uuid.New().String(),
		OrderID:        uuid.New().String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
	}
	actstore := &actteststore.ActivityTestRepo{
		CreateResult: struct {
			ID  *uuid.UUID
			Err error
		}{
			ID:  &activityID,
			Err: nil,
		},
	}
	// TODO: complete here
	acttypestore := &acttypeteststore.ActivityTypeTestRepo{}
	svr := api.NewActivityServer(
		actstore,
		acttypestore,
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordergrpc.ExistOrderReply
			}{
				Err: nil,
				Reply: ordergrpc.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()
	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.From)

	// act
	_, err := svr.Create(ctx, &bookmastergrpc.CreateActivityRequest{
		ActTypeID: activity.ActivityTypeID,
		Period: &bookmastergrpc.Interval{
			From: from,
			To:   to,
		},
	})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking create with non-existing activitytype id provided")
	}

	statusErr := status.Convert(err)
	if statusErr == nil {
		t.Fatalf("provided error is not a status.Status error: %v", err)
	}

	if statusErr.Code() != codes.NotFound {
		t.Errorf("provided error do not present the NotFound code as expected, but instead presents %s", statusErr.Code().String())
	}
}
