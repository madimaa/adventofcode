package day01

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day01.txt")
	leftNumbers := make([]int, 0)
	rightNumbers := make([]int, 0)

	for _, line := range input {
		left, _ := strconv.Atoi(strings.Split(line, "   ")[0])
		right, _ := strconv.Atoi(strings.Split(line, "   ")[1])
		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)

	solution := 0.0
	for i := 0; i < len(leftNumbers); i++ {
		solution += math.Abs(float64(leftNumbers[i]) - float64(rightNumbers[i]))
	}

	log.SetFlags(0)
	log.Printf("%0.f", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day01.txt")
	leftNumbers := make([]int, 0)
	rightNumbers := make([]int, 0)

	for _, line := range input {
		left, _ := strconv.Atoi(strings.Split(line, "   ")[0])
		right, _ := strconv.Atoi(strings.Split(line, "   ")[1])
		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	rightMap := make(map[int]int)
	for _, num := range rightNumbers {
		if _, ok := rightMap[num]; !ok {
			rightMap[num] = 1
		} else {
			rightMap[num]++
		}
	}

	solution := 0
	for _, num := range leftNumbers {
		if _, ok := rightMap[num]; ok {
			solution += num * rightMap[num]
		}
	}

	log.SetFlags(0)
	log.Print(solution)
}
