package conferrs

import "fmt"

// envKeyNotFoundError represent an Not-Found error
type envKeyNotFoundError struct {
	key string
}

func (e *envKeyNotFoundError) Error() string {
	return fmt.Sprintf(`Enviroment variable "%s" not found`, e.key)
}

//NewEnvKeyNotFoundError returns an EnvKeyNotFoundError
func NewEnvKeyNotFoundError(key string) error {
	return &envKeyNotFoundError{key: key}
}
