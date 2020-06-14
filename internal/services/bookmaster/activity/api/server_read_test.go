package api_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	actteststore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage/teststore"
	acttypestorage "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	acttypeteststore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage/teststore"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestReadHappyPath(t *testing.T) {
	// arrange
	orderId, actTypeId, actId := uuid.New(), uuid.New(), uuid.New()

	activity := storage.Activity{
		ID: actId.String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
		OrderID:        orderId.String(),
		ActivityTypeID: actTypeId.String(),
	}
	activityType := bookmastergrpc.ActivityType{
		Id:   actTypeId.String(),
		Code: 1,
		Name: "Test ActivityType",
	}
	order := ordergrpc.Order{
		Id:          orderId.String(),
		Code:        "Order code",
		Description: "Order description",
		Name:        "order name",
	}

	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.To)
	expectedActivity := bookmastergrpc.Activity{
		Id:      activity.ID,
		ActType: &activityType,
		Order:   &order,
		Period: &bookmastergrpc.Interval{
			From: from,
			To:   to,
		},
	}

	actstore := &actteststore.ActivityTestRepo{
		ReadResult: struct {
			Activity *storage.Activity
			Err      error
		}{
			Activity: &activity,
			Err:      nil,
		},
	}
	acttypestore := &acttypeteststore.ActivityTypeTestRepo{
		ReadResult: struct {
			ActivityType *acttypestorage.ActivityType
			Err          error
		}{
			ActivityType: &acttypestorage.ActivityType{
				ID:         activityType.Id,
				Code:       activityType.Code,
				Name:       activityType.Name,
				NeedsOrder: activityType.NeedsOrder,
			},
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		actstore,
		acttypestore,
		&orderTestClient{
			ReadResult: struct {
				Err   error
				Reply ordergrpc.ReadOrderReply
			}{
				Err: nil,
				Reply: ordergrpc.ReadOrderReply{
					Order: &order,
				},
			},
		},
		1*time.Second,
	)

	ctx := context.Background()

	// act
	reply, err := svr.Read(ctx, &bookmastergrpc.ReadActivityRequest{
		Id: activity.ID,
	})
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
	act := reply.GetActivity()

	// assert
	if act == nil {
		t.Fatalf("Returned activity is empty")
	}

	assertEqActivity(t, &expectedActivity, act)
	assertEqActivityType(t, &activityType, act.GetActType())
	assertEqOrder(t, &order, act.GetOrder())
}

func TestRead_NoActivityTypeResponse(t *testing.T) {
	// arrange
	orderId, actTypeId, actId := uuid.New(), uuid.New(), uuid.New()

	activity := storage.Activity{
		ID: actId.String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
		OrderID:        orderId.String(),
		ActivityTypeID: actTypeId.String(),
	}
	activityType := bookmastergrpc.ActivityType{
		Id: actTypeId.String(),
	}
	order := ordergrpc.Order{
		Id:          orderId.String(),
		Code:        "Order code",
		Description: "Order description",
		Name:        "order name",
	}

	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.To)
	expectedActivity := bookmastergrpc.Activity{
		Id:      activity.ID,
		ActType: &activityType,
		Order:   &order,
		Period: &bookmastergrpc.Interval{
			From: from,
			To:   to,
		},
	}
	actstore := &actteststore.ActivityTestRepo{
		ReadResult: struct {
			Activity *storage.Activity
			Err      error
		}{
			Activity: &activity,
			Err:      nil,
		},
	}
	acttypestore := &acttypeteststore.ActivityTypeTestRepo{
		ReadResult: struct {
			ActivityType *acttypestorage.ActivityType
			Err          error
		}{
			ActivityType: nil,
			Err:          status.Error(codes.NotFound, "activity type not found"),
		},
	}
	svr := api.NewActivityServer(
		actstore,
		acttypestore,
		&orderTestClient{
			ReadResult: struct {
				Err   error
				Reply ordergrpc.ReadOrderReply
			}{
				Err: nil,
				Reply: ordergrpc.ReadOrderReply{
					Order: &order,
				},
			},
		},
		1*time.Second,
	)

	ctx := context.Background()

	// act
	reply, err := svr.Read(ctx, &bookmastergrpc.ReadActivityRequest{
		Id: activity.ID,
	})
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
	act := reply.GetActivity()

	// assert
	if act == nil {
		t.Fatalf("Returned activity is empty")
	}

	assertEqActivity(t, &expectedActivity, act)
	assertEqActivityType(t, &activityType, act.GetActType())
	assertEqOrder(t, &order, act.GetOrder())
}

func TestRead_NoOrderResponse(t *testing.T) {
	// arrange
	orderId, actTypeId, actId := uuid.New(), uuid.New(), uuid.New()

	activityType := bookmastergrpc.ActivityType{
		Id:   actTypeId.String(),
		Code: 1,
		Name: "Test ActivityType",
	}
	order := ordergrpc.Order{
		Id: orderId.String(),
	}
	activity := storage.Activity{
		ID: actId.String(),
		Period: storage.Interval{
			From: time.Now(),
			To:   time.Now().Add(10 * 24 * time.Hour),
		},
		OrderID:        orderId.String(),
		ActivityTypeID: actTypeId.String(),
	}

	from, _ := ptypes.TimestampProto(activity.Period.From)
	to, _ := ptypes.TimestampProto(activity.Period.To)
	expectedActivity := bookmastergrpc.Activity{
		Id:      activity.ID,
		ActType: &activityType,
		Order:   &order,
		Period: &bookmastergrpc.Interval{
			From: from,
			To:   to,
		},
	}

	actstore := &actteststore.ActivityTestRepo{
		ReadResult: struct {
			Activity *storage.Activity
			Err      error
		}{
			Activity: &activity,
			Err:      nil,
		},
	}
	acttypestore := &acttypeteststore.ActivityTypeTestRepo{
		ReadResult: struct {
			ActivityType *acttypestorage.ActivityType
			Err          error
		}{
			ActivityType: &acttypestorage.ActivityType{
				ID:         activityType.Id,
				Code:       activityType.Code,
				Name:       activityType.Name,
				NeedsOrder: activityType.NeedsOrder,
			},
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		actstore,
		acttypestore,
		&orderTestClient{
			ReadResult: struct {
				Err   error
				Reply ordergrpc.ReadOrderReply
			}{
				Err:   status.Errorf(codes.Internal, "Internal error test"),
				Reply: ordergrpc.ReadOrderReply{Order: nil},
			},
		},
		1*time.Second,
	)

	ctx := context.Background()

	// act
	reply, err := svr.Read(ctx, &bookmastergrpc.ReadActivityRequest{
		Id: activity.ID,
	})
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}
	act := reply.GetActivity()

	// assert
	if act == nil {
		t.Fatalf("Returned activity is empty")
	}

	assertEqActivity(t, &expectedActivity, act)
	assertEqActivityType(t, &activityType, act.GetActType())
	assertEqOrder(t, &order, act.GetOrder())
}

func isInterfaceNil(c interface{}) bool {
	return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())
}

func assertNil(t *testing.T, expected, provided interface{}) bool {
	expNil, provNil := isInterfaceNil(expected), isInterfaceNil(provided)

	if expNil && provNil {
		return false
	}

	if expNil || provNil {
		if !expNil {
			t.Errorf("provided value is nil while expected is not")
		} else {
			t.Errorf("expected value is nil while provided is not")
		}
		return false
	}
	return true
}

func assertEqActivity(t *testing.T, expected, provided *bookmastergrpc.Activity) {
	if proceed := assertNil(t, expected, provided); !proceed {
		return
	}

	if expected.Id != provided.Id {
		t.Errorf("expected activity Id is %v while provided is %v", expected.Id, provided.Id)
	}

	if expected.Period != nil && provided.Period == nil {
		t.Errorf("expected activity Period is not nil while provided is")
	} else if expected.Period == nil && provided.Period != nil {
		t.Errorf("expected activity Period is nil while provided is not")
	} else {
		if expected.Period.From.Seconds != provided.Period.From.Seconds ||
			expected.Period.From.Nanos != provided.Period.From.Nanos {
			t.Errorf("expected activity Period To is %v while provided is %v", expected.Period.From, provided.Period.From)
		}
		if expected.Period.To.Seconds != provided.Period.To.Seconds ||
			expected.Period.To.Nanos != provided.Period.To.Nanos {
			t.Errorf("expected activity Period To is %v while provided is %v", expected.Period.To, provided.Period.To)
		}
	}
}

func assertEqActivityType(t *testing.T, expected, provided *bookmastergrpc.ActivityType) {
	if proceed := assertNil(t, expected, provided); !proceed {
		return
	}

	if expected.Code != provided.Code {
		t.Errorf("expected activity type Code is %v while provided is %v", expected.Code, provided.Code)
	}
	if expected.Id != provided.Id {
		t.Errorf("expected activity type Id is %v while provided is %v", expected.Id, provided.Id)
	}
	if expected.Name != provided.Name {
		t.Errorf("expected activity type Name is %v while provided is %v", expected.Name, provided.Name)
	}
}

func assertEqOrder(t *testing.T, expected, provided *ordergrpc.Order) {
	if proceed := assertNil(t, expected, provided); !proceed {
		return
	}

	if expected.Code != provided.Code {
		t.Errorf("expected order Code is %v while provided is %v", expected.Code, provided.Code)
	}
	if expected.Id != provided.Id {
		t.Errorf("expected order Id is %v while provided is %v", expected.Id, provided.Id)
	}
	if expected.Name != provided.Name {
		t.Errorf("expected order Name is %v while provided is %v", expected.Name, provided.Name)
	}
	if expected.Description != provided.Description {
		t.Errorf("expected order Description is %v while provided is %v", expected.Description, provided.Description)
	}
}
