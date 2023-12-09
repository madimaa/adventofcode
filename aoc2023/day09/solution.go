package day09

import (
	"log"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day09.txt")
	solution := 0
	for _, line := range input {
		nums := make([]int, 0)
		for _, n := range strings.Split(line, " ") {
			nums = append(nums, util.ConvertToInt(n))
		}
		res := extrapolate(nums, last)
		solution += nums[len(nums)-1] + res
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day09.txt")
	solution := 0
	for _, line := range input {
		nums := make([]int, 0)
		for _, n := range strings.Split(line, " ") {
			nums = append(nums, util.ConvertToInt(n))
		}
		res := extrapolate(nums, first)
		solution += nums[0] - res
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func extrapolate(numbers []int, p position) int {
	diffs := make([]int, 0)
	for i := 1; i < len(numbers); i++ {
		diffs = append(diffs, numbers[i]-numbers[i-1])
	}

	zeros := true
	for _, v := range diffs {
		if v != 0 {
			zeros = false
			break
		}
	}

	if zeros {
		return 0
	} else {
		if p == last {
			return diffs[len(diffs)-1] + extrapolate(diffs, p)
		} else if p == first {
			return diffs[0] - extrapolate(diffs, p)
		}
		panic("Never should have come here...")
	}
}

type position int

const (
	first position = 0
	last  position = 1
)
