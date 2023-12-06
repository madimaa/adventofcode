package day04

import (
	"log"
	"regexp"
	"slices"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day04.txt")
	trimmer := regexp.MustCompile("\\s+")
	cards := make([]card, 0)
	for _, line := range input {
		line = trimmer.ReplaceAllString(line, " ")
		a := strings.Split(line, ": ")

		c := card{id: util.ConvertToInt(strings.Split(a[0], " ")[1]), winningNumbers: make([]int, 0), cardNumbers: make([]int, 0)}
		for _, n := range strings.Split(strings.Split(a[1], " | ")[0], " ") {
			c.winningNumbers = append(c.winningNumbers, util.ConvertToInt(n))
		}

		for _, n := range strings.Split(strings.Split(a[1], " | ")[1], " ") {
			c.cardNumbers = append(c.cardNumbers, util.ConvertToInt(n))
		}

		cards = append(cards, c)
	}

	solution := 0
	for _, c := range cards {
		point := 1
		for _, n := range c.cardNumbers {
			if slices.Contains(c.winningNumbers, n) {
				point *= 2
			}
		}

		solution += point / 2
	}
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day04.txt")
	trimmer := regexp.MustCompile("\\s+")
	cards := make([]card, 0)
	for _, line := range input {
		line = trimmer.ReplaceAllString(line, " ")
		a := strings.Split(line, ": ")

		c := card{id: util.ConvertToInt(strings.Split(a[0], " ")[1]), winningNumbers: make([]int, 0), cardNumbers: make([]int, 0), copies: 1}
		for _, n := range strings.Split(strings.Split(a[1], " | ")[0], " ") {
			c.winningNumbers = append(c.winningNumbers, util.ConvertToInt(n))
		}

		for _, n := range strings.Split(strings.Split(a[1], " | ")[1], " ") {
			c.cardNumbers = append(c.cardNumbers, util.ConvertToInt(n))
		}

		cards = append(cards, c)
	}

	for i := 0; i < len(cards); i++ {
		matches := 0
		c := cards[i]
		for _, n := range c.cardNumbers {
			if slices.Contains(c.winningNumbers, n) {
				matches++
			}
		}

		for j := i + 1; j <= i+matches; j++ {
			if j < len(cards) {
				cards[j].copies += cards[i].copies
			}
		}
	}

	solution := 0
	for _, c := range cards {
		solution += c.copies
	}
	log.SetFlags(0)
	log.Printf("%d", solution)
}

type card struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
	copies         int
}
