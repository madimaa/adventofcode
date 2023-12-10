package day10

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
	"github.com/madimaa/adventofcode/aoc2023/util/array2d"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day10.txt")
	start := getPositionOfS(input)
	neighbours := getConnections(input, start)
	path := make([]point, 2)
	path[0] = start
	path[1] = neighbours[0]
	for {
		next := progress(input, path[len(path)-1], path[len(path)-2])
		if next == start {
			break
		} else {
			path = append(path, next)
		}
	}

	solution := len(path) / 2
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day10.txt")
	start := getPositionOfS(input)
	neighbours := getConnections(input, start)
	path := make([]point, 2)
	path[0] = start
	path[1] = neighbours[0]
	for {
		next := progress(input, path[len(path)-1], path[len(path)-2])
		if next == start {
			break
		} else {
			path = append(path, next)
		}
	}

	startRune := getRuneOfS(path[len(path)-1], path[0], path[1])
	startLine := input[start.y]
	startIndex := strings.IndexRune(startLine, 'S')
	startLine = startLine[:startIndex] + string(startRune) + startLine[startIndex+1:]
	input[start.y] = startLine

	xLen, yLen := len(input[0]), len(input)
	world := array2d.Create(xLen, yLen)

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if []rune(input[y])[x] == '.' {
				world.Put(x, y, '.')
			} else if slices.Contains(path, point{x, y}) {
				world.Put(x, y, []rune(input[y])[x])
			} else {
				world.Put(x, y, '.')
			}
		}
	}

	solution := 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if world.Get(x, y).(rune) == '.' {
				edges := 0
				for k := x; k >= 0; k-- {
					r := world.Get(k, y).(rune)
					if r == '|' || r == '7' || r == 'F' {
						edges++
					}
				}
				if edges%2 != 0 {
					solution++
				}
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func getRuneOfS(last, start, next point) rune {
	switch getDirection(last, start) {
	case 'N':
		switch getDirection(start, next) {
		case 'N':
			return '|'
		case 'E':
			return 'F'
		case 'W':
			return '7'
		}
	case 'E':
		switch getDirection(start, next) {
		case 'N':
			return 'J'
		case 'E':
			return '-'
		case 'S':
			return '7'
		}
	case 'S':
		switch getDirection(start, next) {
		case 'E':
			return 'L'
		case 'S':
			return '|'
		case 'W':
			return 'J'
		}
	case 'W':
		switch getDirection(start, next) {
		case 'N':
			return 'L'
		case 'S':
			return 'F'
		case 'W':
			return '-'
		}
	}
	panic("Never should have come here...")
}

func getPositionOfS(input []string) point {
	for i, v := range input {
		if index := strings.IndexRune(v, 'S'); index != -1 {
			return point{x: index, y: i}
		}
	}

	panic("Something is wrong with the input. There is no S in the file.")
}

func getConnections(input []string, p point) []point {
	neighbours := make([]point, 0)
	northX, northY := p.x, p.y-1
	if northY >= 0 {
		v := []rune(input[northY])[northX]
		if v == '|' || v == '7' || v == 'F' {
			neighbours = append(neighbours, point{x: northX, y: northY})
		}
	}

	southX, southY := p.x, p.y+1
	if southY < len(input) {
		v := []rune(input[southY])[southX]
		if v == '|' || v == 'L' || v == 'J' {
			neighbours = append(neighbours, point{x: southX, y: southY})
		}
	}

	westX, westY := p.x-1, p.y
	if westX >= 0 {
		v := []rune(input[westY])[westX]
		if v == '-' || v == 'L' || v == 'F' {
			neighbours = append(neighbours, point{x: westX, y: westY})
		}
	}

	eastX, eastY := p.x+1, p.y
	if eastX < len(input[eastY]) {
		v := []rune(input[eastY])[eastX]
		if v == '-' || v == 'J' || v == '7' {
			neighbours = append(neighbours, point{x: eastX, y: eastY})
		}
	}

	return neighbours
}

func progress(input []string, start, previous point) point {

	switch []rune(input[start.y])[start.x] {
	case '|':
		if start.y-previous.y > 0 {
			return point{x: start.x, y: start.y + 1}
		} else {
			return point{x: start.x, y: start.y - 1}
		}
	case '-':
		if start.x-previous.x > 0 {
			return point{x: start.x + 1, y: start.y}
		} else {
			return point{x: start.x - 1, y: start.y}
		}
	case 'L':
		if start.x == previous.x {
			return point{x: start.x + 1, y: start.y}
		} else {
			return point{x: start.x, y: start.y - 1}
		}
	case 'J':
		if start.x == previous.x {
			return point{x: start.x - 1, y: start.y}
		} else {
			return point{x: start.x, y: start.y - 1}
		}
	case '7':
		if start.x == previous.x {
			return point{x: start.x - 1, y: start.y}
		} else {
			return point{x: start.x, y: start.y + 1}
		}
	case 'F':
		if start.x == previous.x {
			return point{x: start.x + 1, y: start.y}
		} else {
			return point{x: start.x, y: start.y + 1}
		}
	}

	panic("Never should have come here...")
}

func getDirection(previous, actual point) rune {
	if previous.x > actual.x {
		return 'W'
	}
	if previous.x < actual.x {
		return 'E'
	}
	if previous.y > actual.y {
		return 'N'
	}
	if previous.y < actual.y {
		return 'S'
	}
	panic(fmt.Sprintf("Never should have come here... p: %v, a: %v", previous, actual))
}

type point struct {
	x, y int
}
