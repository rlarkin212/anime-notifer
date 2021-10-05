package util

import "os"

func GetEnvVar(key string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return ""
}
