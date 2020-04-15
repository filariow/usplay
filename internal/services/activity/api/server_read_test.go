package api_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
)

func TestReadHappyPath(t *testing.T) {
	// arrange
	ctime := time.Now()
	ptime, _ := ptypes.TimestampProto(ctime)

	activity := storage.Activity{
		ID:           uuid.New(),
		Code:         "Activity Code",
		Description:  "Activity Description",
		Name:         "Activity Name",
		CreationTime: ctime,
	}
	activityType := activitytypecomm.ActivityType{
		Id:   uuid.New().String(),
		Code: 1,
		Name: "Test ActivityType",
	}
	order := ordercomm.Order{
		Id:          uuid.New().String(),
		Code:        "Order code",
		Description: "Order description",
		Name:        "order name",
	}
	expectedActivity := activitycomm.Activity{
		Code:         activity.Code,
		Id:           activity.ID.String(),
		ActType:      &activityType,
		Order:        &order,
		Description:  activity.Description,
		Name:         activity.Name,
		CreationTime: ptime,
	}

	store := &activityTestRepo{
		ReadResult: struct {
			Activity storage.Activity
			Err      error
		}{
			Activity: activity,
			Err:      nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ReadResult: struct {
				Err   error
				Reply activitytypecomm.ReadActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ReadActivityTypeReply{
					ActivityType: &activityType,
				},
			},
		},
		&orderTestClient{
			ReadResult: struct {
				Err   error
				Reply ordercomm.ReadOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ReadOrderReply{
					Order: &order,
				},
			},
		},
		1*time.Second,
	)

	ctx := context.Background()

	// act
	reply, err := svr.Read(ctx, &activitycomm.ReadActivityRequest{
		Id: activity.ID.String(),
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

func assertEqActivity(t *testing.T, expected, provided *activitycomm.Activity) {
	if proceed := assertNil(t, expected, provided); !proceed {
		return
	}

	if expected.Code != provided.Code {
		t.Errorf("expected activity Code is %v while provided is %v", expected.Code, provided.Code)
	}
	if expected.Description != provided.Description {
		t.Errorf("expected activity Description is %v while provided is %v", expected.Description, provided.Description)
	}
	if expected.Id != provided.Id {
		t.Errorf("expected activity Id is %v while provided is %v", expected.Id, provided.Id)
	}
	if expected.Name != provided.Name {
		t.Errorf("expected activity Name is %v while provided is %v", expected.Name, provided.Name)
	}
}

func assertEqActivityType(t *testing.T, expected, provided *activitytypecomm.ActivityType) {
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

func assertEqOrder(t *testing.T, expected, provided *ordercomm.Order) {
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
