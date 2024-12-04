package day04

import (
	"log"

	"github.com/madimaa/adventofcode/aoc2024/util"
	"github.com/madimaa/adventofcode/aoc2024/util/array2d"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day04.txt")
	xLen, yLen := len(input[0]), len(input)
	puzzle := array2d.CreateEmpty(xLen, yLen, ' ')
	for y, line := range input {
		for x, r := range line {
			puzzle.Put(x, y, r)
		}
	}

	solution := 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if puzzle.Get(x, y).(rune) == 'X' {
				result := getAngles(puzzle, x, y, true)
				for _, res := range result {
					if res == "XMAS" {
						solution++
					}
				}
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day04.txt")
	xLen, yLen := len(input[0]), len(input)
	puzzle := array2d.CreateEmpty(xLen, yLen, ' ')
	for y, line := range input {
		for x, r := range line {
			puzzle.Put(x, y, r)
		}
	}

	solution := 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if puzzle.Get(x, y).(rune) == 'A' {
				result := getAngles(puzzle, x, y, false)
				match := 0
				for _, res := range result {
					if res == "MAS" {
						match++
					}
				}
				if match == 2 {
					solution++
				}
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

// Matrix search algorithm 😉
func getAngles(puzzle *array2d.Array2D, x, y int, p1 bool) []string {
	result := make([]string, 0)

	if p1 {
		result = append(result, string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x+1, y).(rune))+string(puzzle.Get(x+2, y).(rune))+string(puzzle.Get(x+3, y).(rune)))
		result = append(result, string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x-1, y).(rune))+string(puzzle.Get(x-2, y).(rune))+string(puzzle.Get(x-3, y).(rune)))
		result = append(result, string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x, y+1).(rune))+string(puzzle.Get(x, y+2).(rune))+string(puzzle.Get(x, y+3).(rune)))
		result = append(result, string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x, y-1).(rune))+string(puzzle.Get(x, y-2).(rune))+string(puzzle.Get(x, y-3).(rune)))
		result = append(result, string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x+1, y+1).(rune))+string(puzzle.Get(x+2, y+2).(rune))+string(puzzle.Get(x+3, y+3).(rune)))
		result = append(result, string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x-1, y-1).(rune))+string(puzzle.Get(x-2, y-2).(rune))+string(puzzle.Get(x-3, y-3).(rune)))
		result = append(result, string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x-1, y+1).(rune))+string(puzzle.Get(x-2, y+2).(rune))+string(puzzle.Get(x-3, y+3).(rune)))
		result = append(result, string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x+1, y-1).(rune))+string(puzzle.Get(x+2, y-2).(rune))+string(puzzle.Get(x+3, y-3).(rune)))
	} else {
		result = append(result, string(puzzle.Get(x-1, y-1).(rune))+string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x+1, y+1).(rune)))
		result = append(result, string(puzzle.Get(x-1, y+1).(rune))+string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x+1, y-1).(rune)))
		result = append(result, string(puzzle.Get(x+1, y-1).(rune))+string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x-1, y+1).(rune)))
		result = append(result, string(puzzle.Get(x+1, y+1).(rune))+string(puzzle.Get(x, y).(rune))+string(puzzle.Get(x-1, y-1).(rune)))
	}

	return result
}
