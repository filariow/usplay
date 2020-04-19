package storage

import (
	"context"

	"github.com/google/uuid"
)

// Repository Repository interface
type Repository interface {
	Create(context.Context, ActivityType) (*uuid.UUID, error)
	Exist(context.Context, uuid.UUID) (*bool, error)
	Read(context.Context, uuid.UUID) (*ActivityType, error)
	Update(context.Context, ActivityType) error
	Delete(context.Context, uuid.UUID) error
	List(context.Context, []uuid.UUID) (ActivityTypes, error)
}
