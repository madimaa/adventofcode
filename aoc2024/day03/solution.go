package day03

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/madimaa/adventofcode/aoc2024/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day03.txt")
	c := regexp.MustCompile(`(mul\(\d+,\d+\))`)
	solution := 0
	for _, line := range input {
		groups := c.FindAllString(line, -1)
		for _, match := range groups {
			s := match[4 : len(match)-1]
			a, _ := strconv.Atoi(strings.Split(s, ",")[0])
			b, _ := strconv.Atoi(strings.Split(s, ",")[1])
			solution += a * b
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day03.txt")
	c := regexp.MustCompile(`((mul\(\d+,\d+\))|(do\(\))|(don't\(\)))`)
	solution := 0
	do := true
	for _, line := range input {
		groups := c.FindAllString(line, -1)
		for _, match := range groups {
			if len(match) == 0 {
				continue
			}

			if match == "do()" {
				do = true
			} else if match == "don't()" {
				do = false
			} else {
				if do {
					s := match[4 : len(match)-1]
					a, _ := strconv.Atoi(strings.Split(s, ",")[0])
					b, _ := strconv.Atoi(strings.Split(s, ",")[1])
					solution += a * b
				}
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}
