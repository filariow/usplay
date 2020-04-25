package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_UpdateHappyPath(t *testing.T) {
	// arrange
	activity := storage.Activity{
		ID: uuid.New().String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
	}
	store := &activityTestRepo{
		UpdateResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ExistResult: struct {
				Err   error
				Reply activitytypecomm.ExistActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ExistActivityTypeReply{
					Exists: true,
				},
			},
		},
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordercomm.ExistOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()

	// act
	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.To)
	_, err := svr.Update(ctx, &activitycomm.UpdateActivityRequest{
		Id:        activity.ID,
		OrderID:   uuid.New().String(),
		ActTypeID: uuid.New().String(),
		Period: &activitycomm.Interval{
			From: from,
			To:   to,
		},
	})

	// assert
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
}

func Test_UpdateInvalidActivityID(t *testing.T) {
	// arrange
	activity := storage.Activity{
		ActivityTypeID: uuid.New().String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
	}
	store := &activityTestRepo{
		UpdateResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ExistResult: struct {
				Err   error
				Reply activitytypecomm.ExistActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ExistActivityTypeReply{
					Exists: true,
				},
			},
		},
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordercomm.ExistOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()

	// act
	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.To)
	_, err := svr.Update(ctx, &activitycomm.UpdateActivityRequest{
		Id:        "",
		ActTypeID: activity.ActivityTypeID,
		Period: &activitycomm.Interval{
			From: from,
			To:   to,
		},
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

func Test_UpdateInvalidActivityTypeID(t *testing.T) {
	// arrange
	activity := storage.Activity{
		ID: uuid.New().String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
	}
	store := &activityTestRepo{
		UpdateResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ExistResult: struct {
				Err   error
				Reply activitytypecomm.ExistActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ExistActivityTypeReply{
					Exists: true,
				},
			},
		},
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordercomm.ExistOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()

	// act
	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.To)
	_, err := svr.Update(ctx, &activitycomm.UpdateActivityRequest{
		Id:        activity.ID,
		ActTypeID: "",
		OrderID:   uuid.New().String(),
		Period: &activitycomm.Interval{
			From: from,
			To:   to,
		},
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

func Test_UpdateNotExistingActivityTypeID(t *testing.T) {
	// arrange
	activity := storage.Activity{
		ID:             uuid.New().String(),
		ActivityTypeID: uuid.New().String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
	}
	store := &activityTestRepo{
		UpdateResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ExistResult: struct {
				Err   error
				Reply activitytypecomm.ExistActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ExistActivityTypeReply{
					Exists: false,
				},
			},
		},
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordercomm.ExistOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()

	// act
	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.To)
	_, err := svr.Update(ctx, &activitycomm.UpdateActivityRequest{
		Id:        activity.ID,
		ActTypeID: activity.ActivityTypeID,
		OrderID:   uuid.New().String(),
		Period: &activitycomm.Interval{
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

func Test_UpdateInvalidOrderID(t *testing.T) {
	// arrange
	activity := storage.Activity{
		ID: uuid.New().String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
	}
	store := &activityTestRepo{
		UpdateResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ExistResult: struct {
				Err   error
				Reply activitytypecomm.ExistActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ExistActivityTypeReply{
					Exists: true,
				},
			},
		},
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordercomm.ExistOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()

	// act
	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.To)
	_, err := svr.Update(ctx, &activitycomm.UpdateActivityRequest{
		Id:        activity.ID,
		ActTypeID: uuid.New().String(),
		OrderID:   "",
		Period: &activitycomm.Interval{
			From: from,
			To:   to,
		},
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

func Test_UpdateNotExistingOrderID(t *testing.T) {
	// arrange
	activity := storage.Activity{
		ID:             uuid.New().String(),
		ActivityTypeID: uuid.New().String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
	}
	store := &activityTestRepo{
		UpdateResult: struct {
			Err error
		}{
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ExistResult: struct {
				Err   error
				Reply activitytypecomm.ExistActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ExistActivityTypeReply{
					Exists: true,
				},
			},
		},
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordercomm.ExistOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ExistOrderReply{
					Exists: false,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()

	// act
	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.To)
	_, err := svr.Update(ctx, &activitycomm.UpdateActivityRequest{
		Id:        activity.ID,
		ActTypeID: activity.ActivityTypeID,
		OrderID:   uuid.New().String(),
		Period: &activitycomm.Interval{
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
