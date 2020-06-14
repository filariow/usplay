package teststore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	"github.com/google/uuid"
)

// ActivityTestRepo mock test repository
type ActivityTestRepo struct {
	CreateResult struct {
		ID  *uuid.UUID
		Err error
	}
	DeleteResult struct {
		Err error
	}
	ExistResult struct {
		Exists *bool
		Err    error
	}
	ReadResult struct {
		Activity *storage.Activity
		Err      error
	}
	ListResult struct {
		Activities []storage.Activity
		Err        error
	}
	ListInIntervalResult struct {
		Activities []storage.Activity
		Err        error
	}
	UpdateResult struct {
		Err error
	}
}

// Create ...
func (r *ActivityTestRepo) Create(context.Context, storage.Activity) (*uuid.UUID, error) {
	return r.CreateResult.ID, r.CreateResult.Err
}

// Read ...
func (r *ActivityTestRepo) Read(ctx context.Context, id uuid.UUID) (*storage.Activity, error) {
	activity := r.ReadResult.Activity
	activity.ID = id.String()
	return activity, r.ReadResult.Err
}

// Exist ....
func (r *ActivityTestRepo) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	return r.ExistResult.Exists, r.ReadResult.Err
}

// Update ...
func (r *ActivityTestRepo) Update(context.Context, storage.Activity) error {
	return r.UpdateResult.Err
}

// Delete ...
func (r *ActivityTestRepo) Delete(context.Context, uuid.UUID) error {
	return r.DeleteResult.Err
}

// List ...
func (r *ActivityTestRepo) List(context.Context, []uuid.UUID) (storage.Activities, error) {
	return r.ListResult.Activities, r.ListResult.Err
}

// ListInInterval ...
func (r *ActivityTestRepo) ListInInterval(context.Context, storage.Interval) (storage.Activities, error) {
	return r.ListInIntervalResult.Activities, r.ListInIntervalResult.Err
}
