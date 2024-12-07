package day06

import (
	"fmt"
	"image"
	"log"

	"github.com/madimaa/adventofcode/aoc2024/util"
	"github.com/madimaa/adventofcode/aoc2024/util/array2d"
)

var directions = []image.Point{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day06.txt")
	xLen, yLen := len(input[0]), len(input)
	mep := array2d.CreateEmpty(xLen, yLen, '×')
	guardPos := image.Point{}
	for y, line := range input {
		for x, r := range line {
			mep.Put(x, y, r)
			if r == '^' {
				guardPos.X, guardPos.Y = x, y
			}
		}
	}

	solution := len(getVisited(mep, guardPos))

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day06.txt")
	xLen, yLen := len(input[0]), len(input)
	mep := array2d.CreateEmpty(xLen, yLen, '×')
	guardPosOriginal := image.Point{}
	for y, line := range input {
		for x, r := range line {
			mep.Put(x, y, r)
			if r == '^' {
				guardPosOriginal.X, guardPosOriginal.Y = x, y
			}
		}
	}

	solution := 0
	for k := range getVisited(mep, guardPosOriginal) {
		if k == guardPosOriginal {
			continue
		}

		mepCopy := mep.Copy()
		mepCopy.Put(k.X, k.Y, '#')
		visited := make(map[string]bool)
		guardPos := guardPosOriginal
		direction := 0
		for {
			nextPos := guardPos.Add(directions[direction])
			next := mepCopy.Get(nextPos.X, nextPos.Y)
			if next == '×' {
				break
			}
			if next == '#' {
				if _, ok := visited[fmt.Sprintf("%v %v", nextPos, guardPos)]; !ok {
					visited[fmt.Sprintf("%v %v", nextPos, guardPos)] = true
				} else {
					solution++
					break
				}
				direction = (direction + 1) % 4
			} else {
				guardPos = nextPos
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func getVisited(mep *array2d.Array2D, guardPos image.Point) map[image.Point]bool {
	result := make(map[image.Point]bool)
	direction := 0
	for {
		result[guardPos] = true
		nextPos := guardPos.Add(directions[direction])
		next := mep.Get(nextPos.X, nextPos.Y)
		if next == '×' {
			break
		} else if next == '#' {
			direction = (direction + 1) % 4
		} else {
			guardPos = nextPos
		}
	}

	return result
}
