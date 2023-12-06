package util

import (
	"bufio"
	"os"
)

// OpenFile - Open file from path, return file content in string array/slice
func OpenFile(path string) []string {
	file, err := os.Open(path)
	PanicOnError(err)

	scanner := bufio.NewScanner(file)
	fileContent := make([]string, 0)
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	LogOnError(scanner.Err())
	LogOnError(file.Close())

	return fileContent
}
