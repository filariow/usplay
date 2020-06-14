package teststore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/google/uuid"
)

// ActivityTypeTestRepo mock repository
type ActivityTypeTestRepo struct {
	CreateResult struct {
		ID  *uuid.UUID
		Err error
	}
	DeleteResult struct {
		Err error
	}
	ReadResult struct {
		ActivityType *storage.ActivityType
		Err          error
	}
	ExistResult struct {
		Result *bool
		Err    error
	}
	ListResult struct {
		Activities []storage.ActivityType
		Err        error
	}
	UpdateResult struct {
		Err error
	}
}

// New ...
func New() storage.Repository {
	return &ActivityTypeTestRepo{}
}

// Create ...
func (r *ActivityTypeTestRepo) Create(context.Context, storage.ActivityType) (*uuid.UUID, error) {
	return r.CreateResult.ID, r.CreateResult.Err
}

// Exist ....
func (r *ActivityTypeTestRepo) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	return r.ExistResult.Result, r.ExistResult.Err
}

// Read ...
func (r *ActivityTypeTestRepo) Read(ctx context.Context, id uuid.UUID) (*storage.ActivityType, error) {
	return r.ReadResult.ActivityType, r.ReadResult.Err
}

// Update ...
func (r *ActivityTypeTestRepo) Update(context.Context, storage.ActivityType) error {
	return r.UpdateResult.Err
}

// Delete ...
func (r *ActivityTypeTestRepo) Delete(context.Context, uuid.UUID) error {
	return r.DeleteResult.Err
}

// List ...
func (r *ActivityTypeTestRepo) List(context.Context, []uuid.UUID) (storage.ActivityTypes, error) {
	return r.ListResult.Activities, r.ListResult.Err
}
