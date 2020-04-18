package repoerrors

import (
	"fmt"

	"github.com/google/uuid"
)

// NotFoundError represent an Not-Found error
type NotFoundError struct {
	ID uuid.UUID
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("instance with id %s not found", e.ID.String())
}
