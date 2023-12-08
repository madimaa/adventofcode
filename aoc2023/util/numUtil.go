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

// originally from: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// originally from: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
