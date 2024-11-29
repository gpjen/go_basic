package helper

import (
	"os"
	"strconv"
)

func GetEnv[T comparable](key string, defaultValue T) T {
	if value, exist := os.LookupEnv(key); exist {
		switch any(defaultValue).(type) {
		case string:
			return any(value).(T)
		case int:
			if intValue, err := strconv.Atoi(value); err == nil {
				return any(intValue).(T)
			}
		}
	}
	return defaultValue
}
