package day21

import (
	"image"
	"log"

	"github.com/madimaa/adventofcode/aoc2023/util"
	"github.com/madimaa/adventofcode/aoc2023/util/array2d"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day21.txt")
	xLen, yLen := len(input[0]), len(input)
	garden := array2d.CreateEmpty(xLen, yLen, 'X')
	startX, startY := -1, -1
	for y := 0; y < yLen; y++ {
		line := []rune(input[y])
		for x := 0; x < xLen; x++ {
			if line[x] == 'S' {
				startX = x
				startY = y
				garden.Put(x, y, '.')
			} else {
				garden.Put(x, y, line[x])
			}
		}
	}

	possiblePoints := make(map[image.Point]bool, 0)
	possiblePoints[image.Point{startX, startY}] = true
	stepsCount := 64
	for i := 0; i < stepsCount; i++ {
		newPoints := make(map[image.Point]bool, 0)
		for k := range possiblePoints {
			result := progress(garden, k)
			for _, v := range result {
				newPoints[v] = true
			}
		}

		possiblePoints = newPoints
	}

	solution := len(possiblePoints)
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day21.txt")
	xLen, yLen := len(input[0]), len(input)
	garden := array2d.CreateEmpty(xLen, yLen, 'X')
	startX, startY := -1, -1
	for y := 0; y < yLen; y++ {
		line := []rune(input[y])
		for x := 0; x < xLen; x++ {
			if line[x] == 'S' {
				startX = x
				startY = y
				garden.Put(x, y, '.')
			} else {
				garden.Put(x, y, line[x])
			}
		}
	}

	polySteps := []int{startX, startX + xLen, startX + xLen*2}
	polyResults := make([]int, 0)

	for _, stepsCount := range polySteps {
		possiblePoints := make(map[image.Point]bool, 0)
		possiblePoints[image.Point{startX, startY}] = true
		for i := 0; i < stepsCount; i++ {
			newPoints := make(map[image.Point]bool, 0)
			for k := range possiblePoints {
				result := progress2(garden, k)
				for _, v := range result {
					newPoints[v] = true
				}
			}

			possiblePoints = newPoints
		}

		polyResults = append(polyResults, len(possiblePoints))
	}

	steps := 26501365
	poly0 := polyResults[0]
	poly1 := polyResults[1] - polyResults[0]
	poly2 := polyResults[2] - polyResults[1]
	solution := poly0 + (poly1 * (steps / xLen)) + ((steps/xLen)*((steps/xLen)-1)/2)*(poly2-poly1)

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func progress(garden *array2d.Array2D, position image.Point) []image.Point {
	result := make([]image.Point, 0)
	directions := []image.Point{
		{-1, 0}, {0, -1}, {0, 1}, {1, 0},
	}

	for _, v := range directions {
		p := v.Add(position)
		if garden.Get(p.X, p.Y).(rune) == '.' {
			result = append(result, p)
		}
	}

	return result
}

func progress2(garden *array2d.Array2D, position image.Point) []image.Point {
	xLen, yLen := garden.GetSize()
	result := make([]image.Point, 0)
	directions := []image.Point{
		{-1, 0}, {0, -1}, {0, 1}, {1, 0},
	}

	for _, v := range directions {
		p := v.Add(position)
		original := p
		if p.X < 0 {
			p.X %= xLen
			p.X += xLen
		}

		if p.Y < 0 {
			p.Y %= yLen
			p.Y += yLen
		}

		if p.X >= xLen {
			p.X %= xLen
		}

		if p.Y >= yLen {
			p.Y %= yLen
		}

		if garden.Get(p.X, p.Y).(rune) == '.' {
			result = append(result, original)
		}
	}

	return result
}
