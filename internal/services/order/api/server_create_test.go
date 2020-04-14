package api_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/order/api"
	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/google/uuid"
)

func Test_CreateHappyPath(t *testing.T) {
	// arrange
	order := storage.Order{
		Code:        "Order Code",
		Description: "Order Description",
		Name:        "Order Name",
	}
	store := &orderTestRepo{
		CreateResult: struct {
			ID  uuid.UUID
			Err error
		}{
			ID:  order.ID,
			Err: nil,
		},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	reply, err := svr.Create(ctx, &ordercomm.CreateOrderRequest{
		Code:        order.Code,
		Description: order.Description,
		Name:        order.Name,
	})

	// assert
	if err != nil {
		t.Fatalf("error invoking create: %v", err)
	}

	if providedId := reply.GetId(); order.ID.String() != providedId {
		t.Errorf("expected id %s provided %s", order.ID.String(), providedId)
	}
}
