package day08

import (
	"image"
	"log"

	"github.com/madimaa/adventofcode/aoc2024/util"
	"github.com/madimaa/adventofcode/aoc2024/util/array2d"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day08.txt")
	xLen, yLen := len(input[0]), len(input)
	mep := array2d.CreateEmpty(xLen, yLen, '¤')
	frequencies := make(map[rune][]image.Point)
	for y, line := range input {
		for x, r := range []rune(line) {
			mep.Put(x, y, r)
			if r != '.' {
				if _, ok := frequencies[r]; !ok {
					frequencies[r] = make([]image.Point, 0)
				}

				frequencies[r] = append(frequencies[r], image.Point{x, y})
			}
		}
	}

	antinodes := make(map[image.Point]bool)
	for _, coords := range frequencies {
		if len(coords) < 2 {
			continue
		}

		for i := 0; i < len(coords)-1; i++ {
			a := coords[i]
			for j := i + 1; j < len(coords); j++ {
				b := coords[j]
				c := a.Sub(b)
				antinodes[a.Add(c)] = true
				antinodes[b.Sub(c)] = true
			}
		}
	}

	solution := 0

	for pos := range antinodes {
		if mep.Get(pos.X, pos.Y).(rune) != '¤' {
			solution++
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day08.txt")
	xLen, yLen := len(input[0]), len(input)
	mep := array2d.CreateEmpty(xLen, yLen, '¤')
	frequencies := make(map[rune][]image.Point)
	for y, line := range input {
		for x, r := range []rune(line) {
			mep.Put(x, y, r)
			if r != '.' {
				if _, ok := frequencies[r]; !ok {
					frequencies[r] = make([]image.Point, 0)
				}

				frequencies[r] = append(frequencies[r], image.Point{x, y})
			}
		}
	}

	antinodes := make(map[image.Point]bool)
	for _, coords := range frequencies {
		if len(coords) < 2 {
			continue
		}

		for i := 0; i < len(coords)-1; i++ {
			a := coords[i]
			antinodes[a] = true
			for j := i + 1; j < len(coords); j++ {
				b := coords[j]
				c := a.Sub(b)
				mul := 1
				for {
					d := a.Add(c.Mul(mul))
					if mep.Get(d.X, d.Y).(rune) != '¤' {
						antinodes[d] = true
						mul++
					} else {
						break
					}
				}

				mul = 1
				for {
					d := a.Sub(c.Mul(mul))
					if mep.Get(d.X, d.Y).(rune) != '¤' {
						antinodes[d] = true
						mul++
					} else {
						break
					}
				}
			}
		}
	}

	solution := len(antinodes)
	log.SetFlags(0)
	log.Printf("%d", solution)
}
