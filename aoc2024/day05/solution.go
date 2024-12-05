package day05

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/madimaa/adventofcode/aoc2024/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day05.txt")
	rules := make(map[string][]string, 0)
	ruleRead := true
	solution := 0
	for _, line := range input {
		if len(line) == 0 {
			ruleRead = false
			continue
		}

		if ruleRead {
			key := line[:2]
			if _, ok := rules[key]; !ok {
				rules[key] = make([]string, 0)
			}

			rules[key] = append(rules[key], line[3:])
		} else {
			order := strings.Split(line, ",")
			if checkOrder(rules, order) {
				n, _ := strconv.Atoi(order[len(order)/2])
				solution += n
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day05.txt")
	rules := make(map[string][]string, 0)
	ruleRead := true
	solution := 0
	for _, line := range input {
		if len(line) == 0 {
			ruleRead = false
			continue
		}

		if ruleRead {
			key := line[:2]
			if _, ok := rules[key]; !ok {
				rules[key] = make([]string, 0)
			}

			rules[key] = append(rules[key], line[3:])
		} else {
			order := strings.Split(line, ",")
			if !checkOrder(rules, order) {
				order := fixOrder(rules, order)
				n, _ := strconv.Atoi(order[len(order)/2])
				solution += n
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func checkOrder(rules map[string][]string, order []string) bool {
	for i, num := range order {
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(rules[num], order[j]) {
				return false
			}
		}
	}

	return true
}

func fixOrder(rules map[string][]string, order []string) []string {
reset:
	for i, num := range order {
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(rules[num], order[j]) {
				order[i], order[j] = order[j], order[i]
				goto reset
			}
		}
	}

	return order
}
