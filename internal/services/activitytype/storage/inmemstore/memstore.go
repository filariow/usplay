package inmemstore

import (
	"context"
	"fmt"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/errs/repoerrs"
	"github.com/google/uuid"
)

// Configuration Configuration for InMemory store
type Configuration struct{}

type memoryStore struct {
	data map[string]storage.ActivityType
}

// New in memory repository for activities
func New() storage.Repository {
	return &memoryStore{
		data: map[string]storage.ActivityType{},
	}
}

// Create store a new ActivityType
func (s *memoryStore) Create(ctx context.Context, a storage.ActivityType) (*uuid.UUID, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	id := uid.String()

	a.ID = id
	s.data[id] = a
	return &uid, nil
}

// Exist checks if an activitytype exists
func (s *memoryStore) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	_, ok := s.data[id.String()]
	return &ok, nil
}

// Read get an ActivityType by id
func (s *memoryStore) Read(ctx context.Context, id uuid.UUID) (*storage.ActivityType, error) {
	act, ok := s.data[id.String()]
	if !ok {
		return nil, fmt.Errorf("ActivityType with id %v not found", id)
	}
	return &act, nil
}

// Update updates an ActivityType in the store if present
func (s *memoryStore) Update(ctx context.Context, a storage.ActivityType) error {
	if _, ok := s.data[a.ID]; !ok {
		return &repoerrs.NotFoundError{ID: a.ID}
	}

	s.data[a.ID] = a
	return nil
}

// Delete removes an ActivityType from the store if present
func (s *memoryStore) Delete(ctx context.Context, id uuid.UUID) error {
	delete(s.data, id.String())
	return nil
}

// List returns all the activities in the store
func (s *memoryStore) List(ctx context.Context, ids []uuid.UUID) (storage.ActivityTypes, error) {
	list := make([]storage.ActivityType, len(ids), len(s.data))
	for i, v := range ids {
		a, ok := s.data[v.String()]
		if !ok {
			return nil, fmt.Errorf("ID %v not found", v)
		}
		list[i] = a
	}
	return list, nil
}
