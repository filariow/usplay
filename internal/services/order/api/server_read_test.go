package api_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/order/api"
	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/google/uuid"
)

func Test_ReadHappyPath(t *testing.T) {
	// arrange
	order := storage.Order{
		ID:          uuid.New().String(),
		Code:        "Order Code",
		Description: "Order Description",
		Name:        "Order Name",
	}
	store := &orderTestRepo{
		ReadResult: struct {
			Order *storage.Order
			Err   error
		}{
			Order: &order,
			Err:   nil,
		},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	reply, err := svr.Read(ctx, &ordergrpc.ReadOrderRequest{
		Id: order.ID,
	})

	// assert
	if err != nil {
		t.Fatalf("error invoking read: %v", err)
	}

	provOrder := reply.GetOrder()
	if provOrder == nil {
		t.Fatalf("Returned order is nil")
	}

	if providedId := provOrder.GetId(); order.ID != providedId {
		t.Errorf(`expected id %s, provided %s`, order.ID, providedId)
	}
	if providedDesc := provOrder.GetDescription(); order.Description != providedDesc {
		t.Errorf(`expected description "%s", provided "%s"`, order.ID, providedDesc)
	}
	if providedName := provOrder.Name; order.Name != providedName {
		t.Errorf(`expected named "%s", provided "%s"`, order.Name, providedName)
	}
}
