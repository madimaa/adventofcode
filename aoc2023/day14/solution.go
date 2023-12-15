package day14

import (
	"log"

	"github.com/madimaa/adventofcode/aoc2023/util"
	"github.com/madimaa/adventofcode/aoc2023/util/array2d"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day14.txt")
	xLen, yLen := len(input[0]), len(input)
	platform := array2d.Create(xLen, yLen)
	for y, line := range input {
		for x, val := range line {
			platform.Put(x, y, val)
		}
	}

	tilt(0, platform)

	solution := 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if platform.Get(x, y) == 'O' {
				solution += yLen - y
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day14.txt")
	xLen, yLen := len(input[0]), len(input)
	platform := array2d.Create(xLen, yLen)
	for y, line := range input {
		for x, val := range line {
			platform.Put(x, y, val)
		}
	}

	cycles := 1000
	solution := 0
	results := make(map[int][]int)
	for i := 1; i <= cycles; i++ {
		tilt(0, platform)
		tilt(1, platform)
		tilt(2, platform)
		tilt(3, platform)

		solution = 0
		for y := 0; y < yLen; y++ {
			for x := 0; x < xLen; x++ {
				if platform.Get(x, y) == 'O' {
					solution += yLen - y
				}
			}
		}

		if _, ok := results[solution]; !ok {
			results[solution] = make([]int, 0)
		}

		results[solution] = append(results[solution], i)
	}

	for k, v := range results {
		if len(v) > 1 {
			diff := v[1] - v[0]
			same := true
			for i := 2; i < len(v); i++ {
				if diff != v[i]-v[i-1] {
					same = false
					break
				}
			}

			if same {
				if (1000000000-v[0])%diff == 0 {
					solution = k
					break
				}
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func tilt(direction int, platform *array2d.Array2D) {
	xLen, yLen := platform.GetSize()
	switch direction {
	case 0: //north
		for x := 0; x < xLen; x++ {
			roundRocks := 0
			for y := yLen - 1; y >= 0; y-- {
				if platform.Get(x, y) == 'O' {
					roundRocks++
					platform.Put(x, y, '.')
				}
				if platform.Get(x, y) == '#' {
					for i := 0; i < roundRocks; i++ {
						platform.Put(x, y+1+i, 'O')
					}
					roundRocks = 0
				} else if y == 0 {
					for i := 0; i < roundRocks; i++ {
						platform.Put(x, y+i, 'O')
					}
					roundRocks = 0
				}
			}
		}
	case 1: //west
		for y := 0; y < yLen; y++ {
			roundRocks := 0
			for x := xLen - 1; x >= 0; x-- {
				if platform.Get(x, y) == 'O' {
					roundRocks++
					platform.Put(x, y, '.')
				}
				if platform.Get(x, y) == '#' {
					for i := 0; i < roundRocks; i++ {
						platform.Put(x+1+i, y, 'O')
					}
					roundRocks = 0
				} else if x == 0 {
					for i := 0; i < roundRocks; i++ {
						platform.Put(x+i, y, 'O')
					}
					roundRocks = 0
				}
			}
		}
	case 2: //south
		for x := 0; x < xLen; x++ {
			roundRocks := 0
			for y := 0; y < yLen; y++ {
				if platform.Get(x, y) == 'O' {
					roundRocks++
					platform.Put(x, y, '.')
				}
				if platform.Get(x, y) == '#' {
					for i := 0; i < roundRocks; i++ {
						platform.Put(x, y-1-i, 'O')
					}
					roundRocks = 0
				} else if y == yLen-1 {
					for i := 0; i < roundRocks; i++ {
						platform.Put(x, y-i, 'O')
					}
					roundRocks = 0
				}
			}
		}
	case 3: //east
		for y := 0; y < yLen; y++ {
			roundRocks := 0
			for x := 0; x < xLen; x++ {
				if platform.Get(x, y) == 'O' {
					roundRocks++
					platform.Put(x, y, '.')
				}
				if platform.Get(x, y) == '#' {
					for i := 0; i < roundRocks; i++ {
						platform.Put(x-1-i, y, 'O')
					}
					roundRocks = 0
				} else if x == xLen-1 {
					for i := 0; i < roundRocks; i++ {
						platform.Put(x-i, y, 'O')
					}
					roundRocks = 0
				}
			}
		}
	}
}
