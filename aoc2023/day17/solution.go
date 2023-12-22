package day17

import (
	"image"
	"log"
	"math"

	"github.com/madimaa/adventofcode/aoc2023/util"
	pq "github.com/madimaa/adventofcode/aoc2023/util/priorityqueue"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day17.txt")
	xLen, yLen := len(input[0]), len(input)
	city := make(map[image.Point]int)
	for y, line := range input {
		for x, v := range line {
			prio := util.ConvertToInt(string(v))
			city[image.Point{x, y}] = prio
		}
	}

	end := image.Point{xLen - 1, yLen - 1}
	solution := solve(city, end, 1, 3)

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day17.txt")
	xLen, yLen := len(input[0]), len(input)
	city := make(map[image.Point]int)
	for y, line := range input {
		for x, v := range line {
			prio := util.ConvertToInt(string(v))
			city[image.Point{x, y}] = prio
		}
	}

	end := image.Point{xLen - 1, yLen - 1}
	solution := solve(city, end, 4, 10)

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func solve(city map[image.Point]int, end image.Point, minSteps, maxSteps int) int {
	queue := make(pq.PriorityQueue[PPoint], 0)
	visited := make(map[PPoint]bool)
	queue.GPush(PPoint{image.Point{0, 0}, image.Point{1, 0}}, 0)
	queue.GPush(PPoint{image.Point{0, 0}, image.Point{0, 1}}, 0)

	for len(queue) > 0 {
		ppoint, priority := queue.GPop()

		if ppoint.Position == end {
			return priority
		}

		if _, ok := visited[ppoint]; ok {
			continue
		}

		visited[ppoint] = true

		for i := -maxSteps; i <= maxSteps; i++ {
			newPoint := ppoint.Position.Add(ppoint.Direction.Mul(i))
			if _, inCity := city[newPoint]; !inCity || i > -minSteps && i < minSteps {
				continue
			}

			cost, sign := 0, int(math.Copysign(1, float64(i)))
			for j := sign; j != i+sign; j += sign {
				cost += city[ppoint.Position.Add(ppoint.Direction.Mul(j))]
			}
			queue.GPush(PPoint{newPoint, image.Point{ppoint.Direction.Y, ppoint.Direction.X}}, priority+cost)
		}
	}

	return -1
}

type PPoint struct {
	Position  image.Point
	Direction image.Point
}
