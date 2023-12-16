package day16

import (
	"log"
	"slices"

	"github.com/madimaa/adventofcode/aoc2023/util"
	"github.com/madimaa/adventofcode/aoc2023/util/array2d"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day16.txt")
	xLen, yLen := len(input[0]), len(input)
	contraption := array2d.CreateEmpty(xLen, yLen, &tile{make([]*tile, 0), '!', 0, 0})
	for y, line := range input {
		for x, v := range line {
			contraption.Put(x, y, &tile{make([]*tile, 0), v, x, y})
		}
	}

	zero := contraption.Get(0, 0).(*tile)
	zero.cameFrom = append(zero.cameFrom, &tile{make([]*tile, 0), '.', -1, -1})
	followLight(1, zero, zero, contraption)

	solution := 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if len(contraption.Get(x, y).(*tile).cameFrom) != 0 {
				solution++
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day16.txt")
	solution := 0
	xLen, yLen := len(input[0]), len(input)
	for xP := 0; xP < xLen; xP++ {
		contraption := array2d.CreateEmpty(xLen, yLen, &tile{make([]*tile, 0), '!', 0, 0})
		for y, line := range input {
			for x, v := range line {
				contraption.Put(x, y, &tile{make([]*tile, 0), v, x, y})
			}
		}

		zero := contraption.Get(xP, 0).(*tile)
		zero.cameFrom = append(zero.cameFrom, &tile{make([]*tile, 0), '.', -1, -1})
		followLight(2, zero, zero, contraption)

		energized := 0
		for y := 0; y < yLen; y++ {
			for x := 0; x < xLen; x++ {
				if len(contraption.Get(x, y).(*tile).cameFrom) != 0 {
					energized++
				}
			}
		}

		if solution < energized {
			solution = energized
		}
	}

	for xP := 0; xP < xLen; xP++ {
		contraption := array2d.CreateEmpty(xLen, yLen, &tile{make([]*tile, 0), '!', 0, 0})
		for y, line := range input {
			for x, v := range line {
				contraption.Put(x, y, &tile{make([]*tile, 0), v, x, y})
			}
		}

		zero := contraption.Get(xP, yLen-1).(*tile)
		zero.cameFrom = append(zero.cameFrom, &tile{make([]*tile, 0), '.', -1, -1})
		followLight(0, zero, zero, contraption)

		energized := 0
		for y := 0; y < yLen; y++ {
			for x := 0; x < xLen; x++ {
				if len(contraption.Get(x, y).(*tile).cameFrom) != 0 {
					energized++
				}
			}
		}

		if solution < energized {
			solution = energized
		}
	}

	for yP := 0; yP < yLen; yP++ {
		contraption := array2d.CreateEmpty(xLen, yLen, &tile{make([]*tile, 0), '!', 0, 0})
		for y, line := range input {
			for x, v := range line {
				contraption.Put(x, y, &tile{make([]*tile, 0), v, x, y})
			}
		}

		zero := contraption.Get(0, yP).(*tile)
		zero.cameFrom = append(zero.cameFrom, &tile{make([]*tile, 0), '.', -1, -1})
		followLight(1, zero, zero, contraption)

		energized := 0
		for y := 0; y < yLen; y++ {
			for x := 0; x < xLen; x++ {
				if len(contraption.Get(x, y).(*tile).cameFrom) != 0 {
					energized++
				}
			}
		}

		if solution < energized {
			solution = energized
		}
	}

	for yP := 0; yP < yLen; yP++ {
		contraption := array2d.CreateEmpty(xLen, yLen, &tile{make([]*tile, 0), '!', 0, 0})
		for y, line := range input {
			for x, v := range line {
				contraption.Put(x, y, &tile{make([]*tile, 0), v, x, y})
			}
		}

		zero := contraption.Get(xLen-1, yP).(*tile)
		zero.cameFrom = append(zero.cameFrom, &tile{make([]*tile, 0), '.', -1, -1})
		followLight(3, zero, zero, contraption)

		energized := 0
		for y := 0; y < yLen; y++ {
			for x := 0; x < xLen; x++ {
				if len(contraption.Get(x, y).(*tile).cameFrom) != 0 {
					energized++
				}
			}
		}

		if solution < energized {
			solution = energized
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func followLight(direction int, original, actual *tile, contraption *array2d.Array2D) {
	//log.Printf("Follow Light: %v", actual)
	if actual.r == '!' {
		return
	}

	xLen, yLen := contraption.GetSize()
	switch direction {
	case 0: //up
		for y := actual.y; y >= 0; y-- {
			t := contraption.Get(actual.x, y).(*tile)
			if slices.Contains(t.cameFrom, original) {
				return
			}
			t.cameFrom = append(t.cameFrom, original)
			switch t.r {
			case '\\':
				followLight(3, t, contraption.Get(actual.x-1, y).(*tile), contraption)
				return
			case '/':
				followLight(1, t, contraption.Get(actual.x+1, y).(*tile), contraption)
				return
			case '-':
				followLight(3, t, contraption.Get(actual.x-1, y).(*tile), contraption)
				followLight(1, t, contraption.Get(actual.x+1, y).(*tile), contraption)
				return
			}
		}
	case 1: //right
		for x := actual.x; x < xLen; x++ {
			t := contraption.Get(x, actual.y).(*tile)
			if slices.Contains(t.cameFrom, original) {
				return
			}
			t.cameFrom = append(t.cameFrom, original)
			switch t.r {
			case '\\':
				followLight(2, t, contraption.Get(x, actual.y+1).(*tile), contraption)
				return
			case '/':
				followLight(0, t, contraption.Get(x, actual.y-1).(*tile), contraption)
				return
			case '|':
				followLight(0, t, contraption.Get(x, actual.y-1).(*tile), contraption)
				followLight(2, t, contraption.Get(x, actual.y+1).(*tile), contraption)
				return
			}
		}
	case 2: //down
		for y := actual.y; y < yLen; y++ {
			t := contraption.Get(actual.x, y).(*tile)
			if slices.Contains(t.cameFrom, original) {
				return
			}
			t.cameFrom = append(t.cameFrom, original)
			switch t.r {
			case '\\':
				followLight(1, t, contraption.Get(actual.x+1, y).(*tile), contraption)
				return
			case '/':
				followLight(3, t, contraption.Get(actual.x-1, y).(*tile), contraption)
				return
			case '-':
				followLight(3, t, contraption.Get(actual.x-1, y).(*tile), contraption)
				followLight(1, t, contraption.Get(actual.x+1, y).(*tile), contraption)
				return
			}
		}
	case 3: //left
		for x := actual.x; x >= 0; x-- {
			t := contraption.Get(x, actual.y).(*tile)
			if slices.Contains(t.cameFrom, original) {
				return
			}
			t.cameFrom = append(t.cameFrom, original)
			switch t.r {
			case '\\':
				followLight(0, t, contraption.Get(x, actual.y-1).(*tile), contraption)
				return
			case '/':
				followLight(2, t, contraption.Get(x, actual.y+1).(*tile), contraption)
				return
			case '|':
				followLight(0, t, contraption.Get(x, actual.y-1).(*tile), contraption)
				followLight(2, t, contraption.Get(x, actual.y+1).(*tile), contraption)
				return
			}
		}
	}
}

type tile struct {
	cameFrom []*tile
	r        rune
	x, y     int
}

func (t1 *tile) tileEqual(t2 *tile) bool {
	if t1.x == t2.x && t1.y == t2.y {
		return true
	} else {
		return false
	}
}
