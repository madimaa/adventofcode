package day02

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/madimaa/adventofcode/aoc2024/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day02.txt")
	solution := 0

	for _, line := range input {
		split := strings.Split(line, " ")
		_, err := checkForError(split)
		if !err {
			solution++
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day02.txt")
	solution := 0

	for _, line := range input {
		split := strings.Split(line, " ")

		index, err := checkForError(split)
		if err {
			errList := make([]bool, 0)
			if index > 0 {
				splitCopy := make([]string, len(split))
				copy(splitCopy, split)
				_, err1 := checkForError(slices.Delete(splitCopy, index-1, index))
				errList = append(errList, err1)
			}
			if index+1 <= len(split) {
				splitCopy := make([]string, len(split))
				copy(splitCopy, split)
				_, err2 := checkForError(slices.Delete(splitCopy, index, index+1))
				errList = append(errList, err2)
			}
			if index+2 <= len(split) {
				splitCopy := make([]string, len(split))
				copy(splitCopy, split)
				_, err3 := checkForError(slices.Delete(splitCopy, index+1, index+2))
				errList = append(errList, err3)
			}
			if slices.Contains(errList, false) {
				solution++
			}
		} else {
			solution++
		}
	}

	log.SetFlags(0)
	log.Print(solution)
}

func checkForError(numbers []string) (int, bool) {
	monoton := 0
	for i := 0; i < len(numbers)-1; i++ {
		num, _ := strconv.Atoi(numbers[i])

		nextNum, _ := strconv.Atoi(numbers[i+1])
		if num == nextNum {
			return i, true
		}

		div := num - nextNum
		if math.Abs(float64(div)) > 3 {
			return i, true
		}
		if div < 0 && monoton <= 0 {
			monoton = -1
		} else if div > 0 && monoton >= 0 {
			monoton = 1
		} else {
			return i, true
		}
	}

	return 0, false
}
