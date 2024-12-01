package day23

import (
	"image"
	"log"
	"slices"

	"github.com/madimaa/adventofcode/aoc2023/util"
	"github.com/madimaa/adventofcode/aoc2023/util/array2d"
)

type _direction struct {
	direction image.Point
	r         rune
}

type _edge struct {
	length int
	point  image.Point
}

var directions []_direction = []_direction{
	{image.Point{-1, 0}, '>'}, {image.Point{0, -1}, 'v'}, {image.Point{0, 1}, '^'}, {image.Point{1, 0}, '<'},
}

var hike *array2d.Array2D

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day23.txt")
	xLen, yLen := len(input[0]), len(input)
	hike = array2d.CreateEmpty(xLen, yLen, '#')
	for y := 0; y < yLen; y++ {
		line := input[y]
		for x := 0; x < xLen; x++ {
			hike.Put(x, y, []rune(line)[x])
		}
	}

	start := image.Point{1, 0}
	end := image.Point{xLen - 2, yLen - 1}
	path := make([]image.Point, 0)
	path = append(path, start)
	path = calculatePath(path, end)

	solution := len(path) - 1 // the length - 1 = steps
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func calculatePath(path []image.Point, end image.Point) []image.Point {
	for {
		actual := path[len(path)-1]
		if actual.Eq(end) {
			return path
		}

		possibleSteps := make([]image.Point, 0)
		for _, d := range directions {
			p := d.direction.Add(actual)
			if hike.Get(p.X, p.Y).(rune) != '#' && !slices.Contains(path, p) {
				possibleSteps = append(possibleSteps, p)
			}
		}

		switch len(possibleSteps) {
		case 1:
			path = append(path, possibleSteps[0])
		case 2, 3:
			length := 0
			var longest []image.Point
			for _, step := range possibleSteps {
				r := hike.Get(step.X, step.Y).(rune)
				for _, dir := range directions {
					if dir.direction.Eq(actual.Sub(step)) && dir.r == r {
						newPath := append(path, step)
						newPath = calculatePath(newPath, end)
						if length < len(newPath) {
							length = len(newPath)
							longest = newPath
						}
					}
				}
			}

			if length != 0 {
				return longest
			} else {
				return path
			}
		}
	}
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day23.txt")
	xLen, yLen := len(input[0]), len(input)
	hike = array2d.CreateEmpty(xLen, yLen, '#')
	for y := 0; y < yLen; y++ {
		line := input[y]
		for x := 0; x < xLen; x++ {
			hike.Put(x, y, []rune(line)[x])
		}
	}

	start := image.Point{1, 0}
	end := image.Point{xLen - 2, yLen - 1}
	intersections := findIntersections(hike)
	pathMap := mapPossiblePaths(hike, start, end, intersections)
	solution := calculateLongestPath(pathMap, start, end, make([]image.Point, 0))

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func findIntersections(hike *array2d.Array2D) []image.Point {
	intersections := make([]image.Point, 0)
	xLen, yLen := hike.GetSize()
	for x := 1; x < xLen-1; x++ {
		for y := 1; y < yLen-1; y++ {
			if hike.Get(x, y).(rune) != '#' {
				pos := image.Point{X: x, Y: y}
				possibleSteps := make([]image.Point, 0)
				for _, dir := range directions {
					next := pos.Add(dir.direction)
					if hike.Get(next.X, next.Y).(rune) != '#' {
						possibleSteps = append(possibleSteps, next)
					}
				}

				if len(possibleSteps) > 2 {
					intersections = append(intersections, pos)
				}
			}
		}
	}

	return intersections
}

func mapPossiblePaths(hike *array2d.Array2D, start, end image.Point, intersections []image.Point) map[image.Point][]_edge {
	pathMap := make(map[image.Point][]_edge)
	for _, p := range intersections {
		if _, ok := pathMap[p]; !ok {
			pathMap[p] = make([]_edge, 0)
		}
		for _, dir := range directions {
			next := p.Add(dir.direction)
			if hike.Get(next.X, next.Y).(rune) != '#' {
				path := make([]image.Point, 0)
				path = append(path, p, next)
				for {
					actual := path[len(path)-1]
					if actual.Eq(end) || actual.Eq(start) || slices.Contains(intersections, actual) {
						if !slices.Contains(pathMap[p], _edge{point: actual, length: len(path) - 1}) {
							pathMap[p] = append(pathMap[p], _edge{point: actual, length: len(path) - 1})
						}
						if _, ok := pathMap[actual]; !ok {
							pathMap[actual] = make([]_edge, 0)
						}

						if !slices.Contains(pathMap[actual], _edge{point: p, length: len(path) - 1}) {
							pathMap[actual] = append(pathMap[actual], _edge{point: p, length: len(path) - 1})
						}
						break
					}

					possibleSteps := make([]image.Point, 0)
					for _, d := range directions {
						p := d.direction.Add(actual)
						if hike.Get(p.X, p.Y).(rune) != '#' && !slices.Contains(path, p) {
							possibleSteps = append(possibleSteps, p)
						}
					}

					if len(possibleSteps) > 0 {
						path = append(path, possibleSteps[0])
					}
				}
			}
		}
	}

	return pathMap
}

func calculateLongestPath(pathMap map[image.Point][]_edge, start, end image.Point, visited []image.Point) int {
	for _, edge := range pathMap[start] {
		if edge.point.Eq(end) {
			return edge.length
		}
	}
	visited = append(visited, start)

	max := 0
	for _, edge := range pathMap[start] {
		if !slices.Contains(visited, edge.point) {
			act := calculateLongestPath(pathMap, edge.point, end, visited)
			if act+edge.length > max {
				max = act + edge.length
			}
		}
	}

	return max
}
