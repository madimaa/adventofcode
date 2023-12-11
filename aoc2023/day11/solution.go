package day11

import (
	"log"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

var expansion int

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day11.txt")

	expansion = 2
	solution := solve(input)

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day11.txt")

	expansion = 1000000
	solution := solve(input)

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func solve(input []string) (result int) {
	galaxies := make([]galaxy, 0)
	neY := make([]bool, len(input))
	neX := make([]bool, len(input[0]))
	for y, v := range input {
		for x, r := range v {
			if r == '#' {
				neY[y] = true
				neX[x] = true
				galaxies = append(galaxies, galaxy{x, y})
			}
		}
	}

	for i := 0; i < len(galaxies)-1; i++ {
		a := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]
			result += getDistance(a, b, neX, neY)
		}
	}

	return
}

func getDistance(a, b galaxy, neX []bool, neY []bool) (result int) {
	x := a.x - b.x
	if x < 0 {
		x *= -1
	}
	y := a.y - b.y
	if y < 0 {
		y *= -1
	}

	result = x + y
	if neX != nil && neY != nil {
		result += getExpandDistance(a, b, neX, neY)
	}

	return result
}

// getExpandDistance calculates the scale of the expansion between a and b
// because initially in getDistance the original empty line is accounted for
// expansion - 1 should be used all over this function
func getExpandDistance(a, b galaxy, neX []bool, neY []bool) (result int) {
	if a.x < b.x {
		for i := a.x; i < b.x; i++ {
			if !neX[i] {
				result += expansion - 1
			}
		}
	} else {
		for i := b.x; i < a.x; i++ {
			if !neX[i] {
				result += expansion - 1
			}
		}
	}

	if a.y < b.y {
		for i := a.y; i < b.y; i++ {
			if !neY[i] {
				result += expansion - 1
			}
		}
	} else {
		for i := b.y; i < a.y; i++ {
			if !neY[i] {
				result += expansion - 1
			}
		}
	}

	return
}

type galaxy struct {
	x, y int
}
