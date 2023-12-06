package util

import (
	"strconv"
)

// ConvertToInt - takes a string input and converts it to integer
func ConvertToInt(input string) int {
	val, err := strconv.Atoi(input)
	LogOnError(err)
	return val
}

// ConvertToInt64 - takes a string input and converts it to integer64
func ConvertToInt64(input string) int64 {
	val, err := strconv.ParseInt(input, 10, 64)
	LogOnError(err)
	return val
}
