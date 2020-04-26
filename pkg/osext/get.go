package osext

import "os"

//GetEnvOrDefault get environment key value or a default if empty
func GetEnvOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
