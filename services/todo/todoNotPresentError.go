package main

import (
	"fmt"

	"github.com/google/uuid"
)

type todoNotPresentError struct {
	id uuid.UUID
}

func (e *todoNotPresentError) Error() string {
	return fmt.Sprintf("Todo with id %s not found", e.id)
}

// IsTodoNotPresentError checks if the provided error is a TodoNotPresentError
func IsTodoNotPresentError(e error) bool {
	_, ok := e.(*todoNotPresentError)
	return ok
}
