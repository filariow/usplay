package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type memoryStore struct {
	data map[uuid.UUID]Report
}

// NewInMemoryStore in memory repository for reports
func NewInMemoryStore() Repository {
	return &memoryStore{
		data: map[uuid.UUID]Report{},
	}
}

// Generate store a new report
func (s *memoryStore) Create(ctx context.Context, a Report) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	a.ID = id
	s.data[id] = a
	return id, nil
}

// Read get an report by id
func (s *memoryStore) Read(ctx context.Context, id uuid.UUID) (Report, error) {
	act, ok := s.data[id]
	if !ok {
		return Report{}, fmt.Errorf("Report with id %v not found", id)
	}
	return act, nil
}

// Update updates an report in the store if present
func (s *memoryStore) Update(ctx context.Context, a Report) (Report, error) {
	act, err := s.Read(ctx, a.ID)
	if err != nil {
		return act, err
	}

	s.data[a.ID] = a
	return act, nil
}

// Delete removes an report from the store if present
func (s *memoryStore) Delete(ctx context.Context, id uuid.UUID) (Report, error) {
	act, err := s.Read(ctx, id)
	if err != nil {
		return act, err
	}

	delete(s.data, id)
	return act, nil
}

// List returns all the reports in the store
func (s *memoryStore) List(ctx context.Context) (Reports, error) {
	list := make([]Report, 0, len(s.data))
	for _, v := range s.data {
		list = append(list, v)
	}
	return list, nil
}
