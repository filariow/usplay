package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type memoryStore struct {
	data map[uuid.UUID]Activity
}

// NewInMemoryStore in memory repository for activities
func NewInMemoryStore() Repository {
	return &memoryStore{
		data: map[uuid.UUID]Activity{},
	}
}

// Create store a new activity
func (s *memoryStore) Create(ctx context.Context, a Activity) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	a.ID = id
	s.data[id] = a
	return id, nil
}

// Read get an activity by id
func (s *memoryStore) Read(ctx context.Context, id uuid.UUID) (Activity, error) {
	act, ok := s.data[id]
	if !ok {
		return Activity{}, fmt.Errorf("Activity with id %v not found", id)
	}
	return act, nil
}

// Update updates an activity in the store if present
func (s *memoryStore) Update(ctx context.Context, a Activity) error {
	s.data[a.ID] = a
	return nil
}

// Delete removes an activity from the store if present
func (s *memoryStore) Delete(ctx context.Context, id uuid.UUID) error {
	delete(s.data, id)
	return nil
}

// List returns all the activities in the store
func (s *memoryStore) List(ctx context.Context) (Activities, error) {
	list := make([]Activity, 0, len(s.data))
	for _, v := range s.data {
		list = append(list, v)
	}
	return list, nil
}
