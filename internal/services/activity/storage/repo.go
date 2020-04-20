package storage

import (
	"context"

	"github.com/google/uuid"
)

// Repository Repository interface
type Repository interface {
	Create(context.Context, Activity) (*uuid.UUID, error)
	Read(context.Context, uuid.UUID) (*Activity, error)
	Exist(context.Context, uuid.UUID) (*bool, error)
	Update(context.Context, Activity) error
	Delete(context.Context, uuid.UUID) error
	List(context.Context, []uuid.UUID) (Activities, error)
}
