package day12

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day12.txt")
	c := regexp.MustCompile(`([.#?]+)\s([0-9,]+)`)
	hotsprings := make([]hotspring, 0)
	for _, line := range input {
		groups := c.FindStringSubmatch(line)
		hot := hotspring{condition: groups[1], parity: make([]int, 0)}
		for _, v := range strings.Split(groups[2], ",") {
			hot.parity = append(hot.parity, util.ConvertToInt(v))
		}

		hotsprings = append(hotsprings, hot)
	}

	ch := make(chan int)
	for _, v := range hotsprings {
		val := v
		go func() {
			ch <- getDestroyedPossibilities(val.condition, val.parity)
		}()
	}

	solution := 0
	for i := 0; i < len(hotsprings); i++ {
		solution += <-ch
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day12.txt")
	c := regexp.MustCompile(`([.#?]+)\s([0-9,]+)`)
	hotsprings := make([]hotspring, 0)
	for _, line := range input {
		groups := c.FindStringSubmatch(line)
		fiveCondition := make([]string, 0)
		for i := 0; i < 5; i++ {
			fiveCondition = append(fiveCondition, groups[1])
		}

		hotsprng := hotspring{condition: strings.Join(fiveCondition, "?"), parity: make([]int, 0)}
		vals := make([]int, 0)
		for _, v := range strings.Split(groups[2], ",") {
			vals = append(vals, util.ConvertToInt(v))
		}

		for i := 0; i < 5; i++ {
			hotsprng.parity = append(hotsprng.parity, vals...)
		}

		hotsprings = append(hotsprings, hotsprng)
	}

	solution := 0
	for _, v := range hotsprings {
		solution += recursive(v.condition, v.parity, make(map[string]int))
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func getDestroyedPossibilities(condition string, parity []int) int {
	possibilities := make(map[string]bool)
	original := []rune(condition)
	copy := make([]rune, len(condition))
	qc := strings.Count(condition, "?")
	permutations := getPermutations(qc)
	for _, v := range permutations {
		val := []rune(v)
		for i := 0; i < len(original); i++ {
			if original[i] == '?' {
				copy[i] = val[0]
				val = val[1:]
			} else {
				copy[i] = original[i]
			}
		}

		if slices.Equal(parity, getDestroyed(string(copy))) {
			possibilities[string(copy)] = true
		}
	}

	permutations = nil
	return len(possibilities)
}

func getDestroyed(input string) []int {
	result := make([]int, 0)
	c := regexp.MustCompile(`#+`)
	groups := c.FindAllString(input, -1)

	for i := 0; i < len(groups); i++ {
		result = append(result, len(groups[i]))
	}

	return result
}

func getPermutations(length int) []string {
	if length == 1 {
		return []string{".", "#"}
	}

	result := make([]string, 0)
	fill := getPermutations(length - 1)
	for _, v := range fill {
		result = append(result, fmt.Sprint(".", v))
		result = append(result, fmt.Sprint("#", v))
	}

	fill = nil
	return result
}

func recursive(condition string, parity []int, cache map[string]int) (result int) {
	if len(parity) == 0 {
		if !strings.Contains(condition, "#") {
			result = 1
		}

		return
	}

	h := hotspring{condition, parity}
	if v, ok := cache[fmt.Sprintf("%v", h)]; ok {
		return v
	}

	act, parity := parity[0], parity[1:]
	for i := 0; i < len(condition)-util.SumSlice(parity)-len(parity)-act+1; i++ {
		if slices.Contains([]rune(condition)[:i], '#') {
			break
		}
		next := i + act
		noDots := !slices.Contains([]rune(condition)[i:next], '.')
		nextRune := ' '

		if next < len(condition) {
			nextRune = []rune(condition)[next]
		}

		if next <= len(condition) && noDots && nextRune != '#' {
			s := ""
			for j := next + 1; j < len(condition); j++ {
				s += string([]rune(condition)[j])
			}

			result += recursive(s, parity, cache)
		}
	}

	cache[fmt.Sprintf("%v", h)] = result
	return
}

type hotspring struct {
	condition string
	parity    []int
}
