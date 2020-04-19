package repoerrs

import "fmt"

// NotFoundError represent an Not-Found error
type NotFoundError struct {
	ID string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("instance with id %s not found", e.ID)
}
