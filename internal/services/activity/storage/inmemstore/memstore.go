package inmemstore

import (
	"context"
	"fmt"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/google/uuid"
)

//Configuration In Memory Store configuration
type Configuration struct{}

type memoryStore struct {
	data map[uuid.UUID]storage.Activity
}

// New in memory repository for activities
func New() storage.Repository {
	return &memoryStore{
		data: map[uuid.UUID]storage.Activity{},
	}
}

// Create store a new activity
func (s *memoryStore) Create(ctx context.Context, a storage.Activity) (*uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	a.ID = id.String()

	s.data[id] = a
	return &id, nil
}

// Exist get an order by id
func (s *memoryStore) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	_, ok := s.data[id]
	return &ok, nil
}

// Read get an activity by id
func (s *memoryStore) Read(ctx context.Context, id uuid.UUID) (*storage.Activity, error) {
	act, ok := s.data[id]
	if !ok {
		return nil, fmt.Errorf("storage.Activity with id %v not found", id)
	}
	return &act, nil
}

// Update updates an activity in the store if present
func (s *memoryStore) Update(ctx context.Context, a storage.Activity) error {
	id, err := uuid.Parse(a.ID)
	if err != nil {
		return fmt.Errorf("error updating activity, invalid id %s", a.ID)
	}

	s.data[id] = a
	return nil
}

// Delete removes an activity from the store if present
func (s *memoryStore) Delete(ctx context.Context, id uuid.UUID) error {
	delete(s.data, id)
	return nil
}

// List returns all the activities in the store
func (s *memoryStore) List(ctx context.Context, ids []uuid.UUID) (storage.Activities, error) {
	list := make([]storage.Activity, len(ids), len(s.data))
	for i, v := range ids {
		a, ok := s.data[v]
		if !ok {
			return nil, fmt.Errorf("ID %v not found", v)
		}
		list[i] = a
	}
	return list, nil
}

func (s *memoryStore) ListInInterval(ctx context.Context, period storage.Interval) (storage.Activities, error) {
	res := storage.Activities{}
	for _, v := range s.data {
		if v.Period.From.After(period.From) &&
			v.Period.From.Before(period.To) {
			res = append(res, v)
		}
	}
	return res, nil
}
