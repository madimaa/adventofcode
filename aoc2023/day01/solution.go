package day01

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day01.txt")
	sum := 0
	for _, line := range input {
		first := ""
		last := ""
		for _, r := range line {
			if isNumber(r) {
				first = string(r)
				break
			}
		}

		lineRunes := []rune(line)
		for i := len(lineRunes) - 1; i >= 0; i-- {
			if isNumber(lineRunes[i]) {
				last = string(lineRunes[i])
				break
			}
		}

		num, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}
	log.SetFlags(0)
	log.Printf("%d", sum)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day01.txt")
	sum := 0
	for _, line := range input {
		first := ""
		last := ""
		lineRunes := []rune(line)
		for i := 0; i < len(lineRunes); i++ {
			if num := getNumber(lineRunes[i], line[:i+1]); num != -1 {
				first = strconv.Itoa(num)
				break
			}
		}

		for i := len(lineRunes) - 1; i >= 0; i-- {
			if num := getNumber(lineRunes[i], line[i:]); num != -1 {
				last = strconv.Itoa(num)
				break
			}
		}

		num, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}

	log.SetFlags(0)
	log.Print(sum)
}

func getNumber(r rune, s string) int {
	if r >= 49 && r <= 57 {
		return int(r) - 48
	}

	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(nums); i++ {
		if strings.Contains(s, nums[i]) {
			return i + 1
		}
	}

	return -1
}

func isNumber(r rune) bool {
	if r >= 49 && r <= 57 {
		return true
	}

	return false
}
