package day09

import (
	"log"

	"github.com/madimaa/adventofcode/aoc2024/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day09.txt")
	digits := make([]int, 0)
	for _, n := range input[0] {
		digits = append(digits, int(n-'0'))
	}

	multiplier := 0
	sliceIndex := 0
	reverseIndex := len(digits) - 1
	lastFileId := len(digits) / 2
	remaining := digits[reverseIndex]
	solution := 0

	evenIndex := func() {
		fileId := sliceIndex / 2
		for i := 0; i < digits[sliceIndex]; i++ {
			solution += multiplier * fileId
			multiplier++
		}
	}

	oddIndex := func() bool {
		df := digits[sliceIndex]
		for df > 0 {
			for remaining > 0 {
				if df == 0 {
					return false
				}
				solution += lastFileId * multiplier
				multiplier++
				df--
				remaining--
			}

			reverseIndex -= 2
			lastFileId--

			if remaining == 0 && sliceIndex >= reverseIndex {
				return true
			}

			remaining = digits[reverseIndex]
		}

		return false
	}

	for {
		if sliceIndex%2 == 1 || sliceIndex == reverseIndex {
			if oddIndex() {
				break
			}
		} else if sliceIndex%2 == 0 {
			evenIndex()
		}
		sliceIndex++
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day09.txt")
	digits := make([]int, 0)
	for _, n := range input[0] {
		digits = append(digits, int(n-'0'))
	}

	type file struct {
		fileId, size int
	}

	type filemap struct {
		files []file
		free  int
	}

	fileIdMap := make(map[int]filemap)
	movedIndices := make(map[int]int)
	for i := len(digits) - 1; i > 0; i -= 2 {
		for dfi := 1; dfi < i; dfi += 2 {
			free := digits[dfi]
			if _, ok := fileIdMap[dfi]; ok {
				free = fileIdMap[dfi].free
			}

			if free >= digits[i] {
				if _, ok := fileIdMap[dfi]; !ok {
					fileIdMap[dfi] = filemap{files: make([]file, 0), free: 0}
				}

				m := fileIdMap[dfi]
				m.files = append(m.files, file{fileId: i / 2, size: digits[i]})
				m.free = free - digits[i]
				fileIdMap[dfi] = m
				movedIndices[i] = digits[i]
				break
			}
		}
	}

	multiplier := 0
	solution := 0
	for sliceIndex, digit := range digits {
		if _, ok := movedIndices[sliceIndex]; !ok && sliceIndex%2 == 0 {
			fileId := sliceIndex / 2
			for i := 0; i < digit; i++ {
				solution += multiplier * fileId
				multiplier++
			}
		} else if sliceIndex%2 == 1 {
			if _, ok := fileIdMap[sliceIndex]; ok {
				for _, file := range fileIdMap[sliceIndex].files {
					for j := 0; j < file.size; j++ {
						solution += multiplier * file.fileId
						multiplier++
					}
				}

				multiplier += fileIdMap[sliceIndex].free
			} else {
				multiplier += digits[sliceIndex]
			}
		} else if _, ok := movedIndices[sliceIndex]; ok {
			multiplier += movedIndices[sliceIndex]
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}
