package api_test

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/google/uuid"
)

// Order Test repository
type orderTestRepo struct {
	CreateResult struct {
		ID  *uuid.UUID
		Err error
	}
	DeleteResult struct {
		Err error
	}
	ExistResult struct {
		Exist *bool
		Err   error
	}
	ReadResult struct {
		Order *storage.Order
		Err   error
	}
	ListResult struct {
		Orders storage.Orders
		Err    error
	}
	UpdateResult struct {
		Err error
	}
}

// Create
func (r *orderTestRepo) Create(context.Context, storage.Order) (*uuid.UUID, error) {
	return r.CreateResult.ID, r.CreateResult.Err
}

// Exist
func (r *orderTestRepo) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	return r.ExistResult.Exist, r.ExistResult.Err
}

// Read
func (r *orderTestRepo) Read(ctx context.Context, id uuid.UUID) (*storage.Order, error) {
	order := r.ReadResult.Order
	order.ID = id.String()
	return order, r.ReadResult.Err
}

// Update
func (r *orderTestRepo) Update(context.Context, storage.Order) error {
	return r.UpdateResult.Err
}

// Delete
func (r *orderTestRepo) Delete(context.Context, uuid.UUID) error {
	return r.DeleteResult.Err
}

// List
func (r *orderTestRepo) List(context.Context, []uuid.UUID) (storage.Orders, error) {
	return r.ListResult.Orders, r.ListResult.Err
}
