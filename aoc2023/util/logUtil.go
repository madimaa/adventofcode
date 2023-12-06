package util

import "log"

// LogOnError - check and log the error
func LogOnError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// PanicOnError - panic on error
func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}
