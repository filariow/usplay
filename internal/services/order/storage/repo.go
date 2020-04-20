package storage

import (
	"context"

	"github.com/google/uuid"
)

// Repository Repository interface
type Repository interface {
	Create(context.Context, Order) (*uuid.UUID, error)
	Read(context.Context, uuid.UUID) (*Order, error)
	Exist(context.Context, uuid.UUID) (*bool, error)
	Update(context.Context, Order) error
	Delete(context.Context, uuid.UUID) error
	List(context.Context, []uuid.UUID) (Orders, error)
}
