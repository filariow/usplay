package storage

import (
	"context"

	"github.com/google/uuid"
)

// Repository Repository interface
type Repository interface {
	Create(context.Context, Activity) (uuid.UUID, error)
	Read(context.Context, uuid.UUID) (Activity, error)
	Update(context.Context, Activity) (Activity, error)
	Delete(context.Context, uuid.UUID) (Activity, error)
	List(context.Context) (Activities, error)
}
