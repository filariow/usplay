package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type memoryStore struct {
	data map[uuid.UUID]Order
}

// NewInMemoryStore in memory repository for activities
func NewInMemoryStore() Repository {
	return &memoryStore{
		data: map[uuid.UUID]Order{},
	}
}

// Create store a new order
func (s *memoryStore) Create(ctx context.Context, a Order) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	a.ID = id
	s.data[id] = a
	return id, nil
}

// Read get an order by id
func (s *memoryStore) Read(ctx context.Context, id uuid.UUID) (Order, error) {
	act, ok := s.data[id]
	if !ok {
		return Order{}, fmt.Errorf("Order with id %v not found", id)
	}
	return act, nil
}

// Update updates an order in the store if present
func (s *memoryStore) Update(ctx context.Context, a Order) (Order, error) {
	act, err := s.Read(ctx, a.ID)
	if err != nil {
		return act, err
	}

	s.data[a.ID] = a
	return act, nil
}

// Delete removes an order from the store if present
func (s *memoryStore) Delete(ctx context.Context, id uuid.UUID) (Order, error) {
	act, err := s.Read(ctx, id)
	if err != nil {
		return act, err
	}

	delete(s.data, id)
	return act, nil
}

// List returns all the activities in the store
func (s *memoryStore) List(ctx context.Context) (Activities, error) {
	list := make([]Order, 0, len(s.data))
	for _, v := range s.data {
		list = append(list, v)
	}
	return list, nil
}
