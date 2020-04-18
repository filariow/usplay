package api_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/order/api"
	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_ExistHappyPathTrue(t *testing.T) {
	// arrange
	order := storage.Order{
		ID:          uuid.New(),
		Code:        "Order Code",
		Description: "Order Description",
		Name:        "Order Name",
	}
	store := &orderTestRepo{
		ExistResult: struct {
			Exist bool
			Err   error
		}{
			Exist: true,
			Err:   nil,
		},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	reply, err := svr.Exist(ctx, &ordercomm.ExistOrderRequest{
		Id: order.ID.String(),
	})

	// assert
	if err != nil {
		t.Fatalf("error invoking exist: %v", err)
	}

	if exist := reply.GetExists(); !exist {
		t.Fatalf("Expected result is exists=true, obtained false")
	}
}

func Test_ExistHappyPathFalse(t *testing.T) {
	// arrange
	order := storage.Order{
		ID:          uuid.New(),
		Code:        "Order Code",
		Description: "Order Description",
		Name:        "Order Name",
	}
	store := &orderTestRepo{
		ExistResult: struct {
			Exist bool
			Err   error
		}{
			Exist: false,
			Err:   nil,
		},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	reply, err := svr.Exist(ctx, &ordercomm.ExistOrderRequest{
		Id: order.ID.String(),
	})

	// assert
	if err != nil {
		t.Fatalf("error invoking exist: %v", err)
	}

	if exist := reply.GetExists(); exist {
		t.Fatalf("Expected result is exists=false, obtained true")
	}
}

func Test_ExistSadPath(t *testing.T) {
	// arrange
	order := storage.Order{
		ID:          uuid.New(),
		Code:        "Order Code",
		Description: "Order Description",
		Name:        "Order Name",
	}
	store := &orderTestRepo{
		ExistResult: struct {
			Exist bool
			Err   error
		}{
			Exist: false,
			Err:   fmt.Errorf("Test Error"),
		},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Exist(ctx, &ordercomm.ExistOrderRequest{
		Id: order.ID.String(),
	})

	// assert
	if err == nil {
		t.Fatalf("Error was expected, but non provided")
	}

	stErr := status.Convert(err)
	if stErr == nil {
		t.Fatalf("Provided error is not of type status.Status")
	}

	if stErr.Code() != codes.NotFound {
		t.Fatalf("Expected error status code %s, obtained %s", codes.NotFound, stErr.Code())
	}
}
