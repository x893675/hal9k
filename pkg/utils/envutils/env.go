package envutils

import (
	"hal9k/pkg/utils/stringutils"
	"os"
)

func GetEnvironment(key, value string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return value
	}
	return val
}

func GetEnvironmentToInt(key, value string) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return stringutils.S(value).DefaultInt(0)
	}
	return stringutils.S(val).DefaultInt(0)
}

func GetEnvironmentToBool(key, value string) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return stringutils.S(value).DefaultBool(false)
	}
	return stringutils.S(val).DefaultBool(false)
}
