package day13

import (
	"log"
	"slices"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day13.txt")
	current := make([]string, 0)
	counter := 0
	solution := 0
	for _, line := range input {
		if len(line) == 0 {
			solution += 100 * findReflection(current)
			solution += findReflection(rotate(current))
			current = make([]string, 0)
			counter = 0
			continue
		}

		current = append(current, line)
		counter++
	}

	solution += 100 * findReflection(current)
	solution += findReflection(rotate(current))

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day13.txt")
	current := make([]string, 0)
	counter := 0
	solution := 0
	for _, line := range input {
		if len(line) == 0 {
			solution += 100 * findSmudge(current)
			solution += findSmudge(rotate(current))

			current = make([]string, 0)
			counter = 0
			continue
		}

		current = append(current, line)
		counter++
	}

	solution += 100 * findSmudge(current)
	solution += findSmudge(rotate(current))

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func findReflectionOld(input []string) (int, int) {
	length := 0
	longest := 0
	for i := 1; i < len(input); i++ {
		if input[i-1] == input[i] {
			l := 1
			reflect := true
			for j := i + 1; j < len(input); j++ {
				if j >= len(input) {
					break
				}

				j2 := 2*i - j - 1
				if j2 < 0 {
					break
				}

				if input[j] != input[j2] {
					reflect = false
					break
				}

				l++
			}

			if reflect {
				if l > length {
					length = l
					longest = i
				}
			}
		}
	}

	return longest, length
}

func findReflection(input []string) int {
	for i := 1; i < len(input); i++ {
		if input[i-1] == input[i] {
			touchy := true

			for j := i + 1; j < len(input); j++ {
				j2 := 2*i - j - 1
				if j2 < 0 {
					break
				}

				if input[j] != input[j2] {
					touchy = false
					break
				}
			}

			if touchy {
				return i
			}
		}
	}

	return 0
}

func rotate(input []string) []string {
	res := make([]string, len(input[0]))
	for _, line := range input {
		for j, v := range line {
			res[j] += string(v)
		}
	}

	return res
}

func findSmudge(input []string) int {
	oldPlace := findReflection(input)
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			a := input[i]
			b := []rune(input[j])
			diffs := 0
			for k, v := range a {
				if v != b[k] {
					diffs++
					if diffs > 1 {
						goto CONT
					}
				}
			}

			if diffs == 1 {
				new := slices.Clone(input)
				new[i] = new[j]
				newPlace := findReflection(new)
				if newPlace != oldPlace {
					return newPlace
				}
			}
		CONT:
		}
	}

	return 0
}

type mirror struct {
}
