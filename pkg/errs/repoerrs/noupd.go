package repoerrs

import "fmt"

// NotUpdatedError represent an Not-Updated error
type NotUpdatedError struct {
	ID string
}

func (e *NotUpdatedError) Error() string {
	return fmt.Sprintf("instance with id %s not updated", e.ID)
}
