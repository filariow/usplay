package api_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/order/api"
	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/google/uuid"
)

func Test_ListHappyPath(t *testing.T) {
	// arrange
	orders := make([]storage.Order, 2)
	for i := 1; i <= 2; i++ {
		order := storage.Order{
			ID:          uuid.New(),
			Code:        fmt.Sprintf("Order Code %v", i),
			Description: fmt.Sprintf("Order Description %v", i),
			Name:        fmt.Sprintf("Order Name %v", i),
		}
		orders[i-1] = order
	}

	store := &orderTestRepo{
		ListResult: struct {
			Orders storage.Orders
			Err    error
		}{
			Orders: orders,
			Err:    nil,
		},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	reply, err := svr.List(ctx, &ordercomm.ListOrdersRequest{})

	// assert
	if err != nil {
		t.Fatalf("error invoking list: %v", err)
	}

	provOrders := reply.GetOrders()
	if provOrders == nil {
		t.Fatalf("Returned orders is nil")
	}

	for i, provOrder := range provOrders {
		order := orders[i]

		if providedId := provOrder.GetId(); order.ID.String() != providedId {
			t.Errorf(`expected id %s, provided %s`, order.ID.String(), providedId)
		}
		if providedDesc := provOrder.GetDescription(); order.Description != providedDesc {
			t.Errorf(`expected description "%s", provided "%s"`, order.ID.String(), providedDesc)
		}
		if providedName := provOrder.Name; order.Name != providedName {
			t.Errorf(`expected named "%s", provided "%s"`, order.Name, providedName)
		}
	}
}
