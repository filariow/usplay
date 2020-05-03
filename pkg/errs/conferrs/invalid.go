package conferrs

import "fmt"

// envKeyInvalidError represent an Invalid EnvVar error
type envKeyInvalidError struct {
	Key   string
	Value string
}

func (e *envKeyInvalidError) Error() string {
	return fmt.Sprintf(`Environment variable "%s" value is invalid: value is %s`, e.Key, e.Value)
}

// NewEnvKeyInvalidError returns an EnvKeyInvalidError
func NewEnvKeyInvalidError(key, value string) error {
	return &envKeyInvalidError{
		Key:   key,
		Value: value,
	}
}
