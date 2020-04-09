package storage

import (
	"context"

	"github.com/google/uuid"
)

// Repository Repository interface
type Repository interface {
	Create(context.Context, Order) (uuid.UUID, error)
	Read(context.Context, uuid.UUID) (Order, error)
	Update(context.Context, Order) (Order, error)
	Delete(context.Context, uuid.UUID) (Order, error)
	List(context.Context) (Activities, error)
}
