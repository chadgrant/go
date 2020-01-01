package infra

import (
	"os"
	"strconv"
)

func GetEnvVar(name string, def string) string {
	if v := os.Getenv(name); len(v) > 0 {
		return v
	}
	return def
}

func GetEnvVarInt(name string, def int) int {
	if v := os.Getenv(name); len(v) > 0 {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}
func GetEnvVarBool(name string, def bool) bool {
	if v := os.Getenv(name); len(v) > 0 {
		if i, err := strconv.ParseBool(v); err == nil {
			return i
		}
	}
	return def
}
