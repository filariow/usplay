package api_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/order/api"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_UpdateHappyPath(t *testing.T) {
	// arrange
	store := &orderTestRepo{
		UpdateResult: struct{ Err error }{Err: nil},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Update(ctx, &ordercomm.UpdateOrderRequest{Id: uuid.New().String()})

	// assert
	if err != nil {
		t.Fatalf("error invoking update: %v", err)
	}
}

func Test_UpdateInvalidId(t *testing.T) {
	// arrange
	store := &orderTestRepo{
		UpdateResult: struct{ Err error }{Err: nil},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Update(ctx, &ordercomm.UpdateOrderRequest{Id: "88888128312319273190"})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking update with empty id, but none received")
	}

	sts := status.Convert(err)
	if sts == nil {
		t.Fatalf("error is not a status.Status error")
	}

	if expected := codes.InvalidArgument; sts.Code() != expected {
		t.Errorf("expected status code %s, provided code %s", expected, sts.Code())
	}
}

func Test_UpdateEmptyId(t *testing.T) {
	// arrange
	store := &orderTestRepo{
		UpdateResult: struct{ Err error }{Err: nil},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Update(ctx, &ordercomm.UpdateOrderRequest{Id: ""})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking update with empty id, but none received")
	}

	sts := status.Convert(err)
	if sts == nil {
		t.Fatalf("error is not a status.Status error")
	}

	if expected := codes.InvalidArgument; sts.Code() != expected {
		t.Errorf("expected status code %s, provided code %s", expected, sts.Code())
	}
}

func Test_UpdateNilId(t *testing.T) {
	// arrange
	store := &orderTestRepo{
		UpdateResult: struct{ Err error }{Err: nil},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Update(ctx, &ordercomm.UpdateOrderRequest{Id: uuid.Nil.String()})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking update with nil id, but none received")
	}

	sts := status.Convert(err)
	if sts == nil {
		t.Fatalf("error is not a status.Status error")
	}

	if expected := codes.InvalidArgument; sts.Code() != expected {
		t.Errorf("expected status code %s, provided code %s", expected, sts.Code())
	}
}
