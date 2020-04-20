package inmemstore

import (
	"context"
	"fmt"

	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/google/uuid"
)

// Configuration Configuration for InMemory store
type Configuration struct{}

type memoryStore struct {
	data map[uuid.UUID]storage.Order
}

// New in memory repository for activities
func New() storage.Repository {
	return &memoryStore{
		data: map[uuid.UUID]storage.Order{},
	}
}

// Create store a new order
func (s *memoryStore) Create(ctx context.Context, a storage.Order) (*uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	a.ID = id.String()
	s.data[id] = a
	return &id, nil
}

// Read get an order by id
func (s *memoryStore) Read(ctx context.Context, id uuid.UUID) (*storage.Order, error) {
	act, ok := s.data[id]
	if !ok {
		return nil, fmt.Errorf("storage.Order with id %v not found", id)
	}
	return &act, nil
}

// Exist get an order by id
func (s *memoryStore) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	_, ok := s.data[id]
	return &ok, nil
}

// Update updates an order in the store if present
func (s *memoryStore) Update(ctx context.Context, a storage.Order) error {
	id, err := uuid.Parse(a.ID)
	if err != nil {
		return fmt.Errorf("error updating order, invalid id %s", a.ID)
	}
	
	if _, err := s.Read(ctx, id); err != nil {
		return err
	}

	s.data[id] = a
	return nil
}

// Delete removes an order from the store if present
func (s *memoryStore) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.Read(ctx, id); err != nil {
		return err
	}

	delete(s.data, id)
	return nil
}

// List returns all the activities in the store
func (s *memoryStore) List(ctx context.Context, ids []uuid.UUID) (storage.Orders, error) {
	list := make([]storage.Order, len(ids), len(s.data))
	for i, v := range ids {
		a, ok := s.data[v]
		if !ok {
			return nil, fmt.Errorf("ID %v not found", v)
		}
		list[i] = a
	}
	return list, nil
}
