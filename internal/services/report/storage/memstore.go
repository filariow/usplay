package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type memoryStore struct {
	data map[string]Report
}

// NewInMemoryStore in memory repository for reports
func NewInMemoryStore() Repository {
	return &memoryStore{
		data: map[string]Report{},
	}
}

// Generate store a new report
func (s *memoryStore) Create(ctx context.Context, a Report) (*uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	a.ID = id.String()
	s.data[a.ID] = a
	return &id, nil
}

// Read get an report by id
func (s *memoryStore) Read(ctx context.Context, id uuid.UUID) (*Report, error) {
	report, ok := s.data[id.String()]
	if !ok {
		return nil, fmt.Errorf("Report with id %v not found", id)
	}
	return &report, nil
}

// Update updates an report in the store if present
func (s *memoryStore) Update(ctx context.Context, a Report) error {
	if _, ok := s.data[a.ID]; !ok {
		return fmt.Errorf("Report with id %v not found", a.ID)
	}

	s.data[a.ID] = a
	return nil
}

// Delete removes an report from the store if present
func (s *memoryStore) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.Read(ctx, id); err != nil {
		return err
	}

	delete(s.data, id.String())
	return nil
}

// List returns all the reports in the store
func (s *memoryStore) List(ctx context.Context) (Reports, error) {
	list := make([]Report, 0, len(s.data))
	for _, v := range s.data {
		list = append(list, v)
	}
	return list, nil
}
