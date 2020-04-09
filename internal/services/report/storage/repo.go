package storage

import (
	"context"

	"github.com/google/uuid"
)

// Repository Repository interface
type Repository interface {
	Create(context.Context, Report) (uuid.UUID, error)
	Read(context.Context, uuid.UUID) (Report, error)
	Update(context.Context, Report) (Report, error)
	Delete(context.Context, uuid.UUID) (Report, error)
	List(context.Context) (Reports, error)
}
