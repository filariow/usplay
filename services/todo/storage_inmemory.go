package main

import (
	"fmt"

	"github.com/google/uuid"
)

// InMemoryRepository In memory repository for todo
type InMemoryRepository struct {
	todos map[uuid.UUID]Todo
}

// NewInMemoryRepository instantiates a new In Memory repository for todos
func NewInMemoryRepository() Repository {
	return &InMemoryRepository{
		todos: make(map[uuid.UUID]Todo),
	}
}

// Create create todo
func (r *InMemoryRepository) Create(t Todo) (Todo, error) {
	t.ID = uuid.New()
	r.todos[t.ID] = t
	return t, nil
}

// Update if it exists updates an entry and returns the previous version, otherwise it creates a new one
func (r *InMemoryRepository) Update(t Todo) (*Todo, error) {
	if t.ID == uuid.Nil {
		return nil, fmt.Errorf("Provided todo must have a valid id: nil id found")
	}

	oldT, _ := r.todos[t.ID]
	r.todos[t.ID] = t
	return &oldT, nil
}

// Delete delete the todo with the given id
func (r *InMemoryRepository) Delete(id uuid.UUID) (*Todo, error) {
	ftodo, ok := r.todos[id]
	if !ok {
		return nil, &todoNotPresentError{id: id}
	}

	delete(r.todos, id)
	return &ftodo, nil
}

// ReadAll read all todos
func (r *InMemoryRepository) ReadAll() ([]Todo, error) {
	todos := []Todo{}
	for _, t := range r.todos {
		todos = append(todos, t)
	}
	return todos, nil
}

// Read returns the todo with the given id
func (r *InMemoryRepository) Read(id uuid.UUID) (*Todo, error) {
	for _, t := range r.todos {
		if t.ID == id {
			res := t
			return &res, nil
		}
	}
	return nil, nil
}
