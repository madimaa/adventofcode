package day03

import (
	"fmt"
	"log"
	"strconv"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day03.txt")
	symbols := make(map[string]string)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if !isNumber([]rune(input[i])[j]) && string(input[i][j]) != "." {
				symbols[fmt.Sprintf("%d %d", i, j)] = string(input[i][j])
			}
		}
	}

	solution := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if isNumber([]rune(input[i])[j]) {
				p := collectNumber(part{adjacentSymbols: make(map[string]string)}, input, symbols, j, i)
				j += len(p.partNumber) - 1
				if p.adjacentSymbol {
					val, err := strconv.Atoi(p.partNumber)
					if err != nil {
						log.Fatal(err)
					}
					solution += val
				}
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day03.txt")
	symbols := make(map[string]string)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if !isNumber([]rune(input[i])[j]) && string(input[i][j]) != "." {
				symbols[fmt.Sprintf("%d %d", i, j)] = string(input[i][j])
			}
		}
	}

	parts := make([]part, 0)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if isNumber([]rune(input[i])[j]) {
				p := collectNumber(part{adjacentSymbols: make(map[string]string)}, input, symbols, j, i)
				j += len(p.partNumber) - 1
				parts = append(parts, p)
			}
		}
	}

	solution := 0
	for k, v := range symbols {
		if v == "*" {
			adjacents := make([]part, 0)
			for _, p := range parts {
				if _, ok := p.adjacentSymbols[k]; ok {
					adjacents = append(adjacents, p)
				}
			}

			if len(adjacents) == 2 {
				a, err := strconv.Atoi(adjacents[0].partNumber)
				if err != nil {
					log.Fatal(err)
				}

				b, err := strconv.Atoi(adjacents[1].partNumber)
				if err != nil {
					log.Fatal(err)
				}

				solution += a * b
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func collectNumber(p part, input []string, symbols map[string]string, x, y int) part {
	line := input[y]
	if isNumber([]rune(line)[x]) {
		p.partNumber += string(line[x])
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				a := y + i
				b := x + j
				if val, ok := symbols[fmt.Sprintf("%d %d", a, b)]; ok {
					p.adjacentSymbol = true
					p.adjacentSymbols[fmt.Sprintf("%d %d", a, b)] = val
				}
			}
		}
	}

	if x+1 < len(line) && isNumber([]rune(line)[x+1]) {
		p = collectNumber(p, input, symbols, x+1, y)
	}

	return p
}

type part struct {
	partNumber      string
	adjacentSymbol  bool
	adjacentSymbols map[string]string
}

func isNumber(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}

	return false
}
