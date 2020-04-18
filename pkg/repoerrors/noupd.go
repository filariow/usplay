package repoerrors

import (
	"fmt"

	"github.com/google/uuid"
)

// NotUpdatedError represent an Not-Updated error
type NotUpdatedError struct {
	ID uuid.UUID
}

func (e *NotUpdatedError) Error() string {
	return fmt.Sprintf("instance with id %s not updated", e.ID.String())
}
