package day07

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

var part2 bool

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day07.txt")
	hands := make([]hand, 0)
	part2 = false
	for _, line := range input {
		a := strings.Split(line, " ")
		hands = append(hands, hand{cards: a[0], bid: util.ConvertToInt(a[1]), tyype: determineHandType(a[0])})
	}

	solution := 0
	slices.SortFunc(hands, compareCards)
	for i, v := range hands {
		solution += (i + 1) * v.bid
	}
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day07.txt")
	hands := make([]hand, 0)
	part2 = true
	for _, line := range input {
		a := strings.Split(line, " ")
		hands = append(hands, hand{cards: a[0], bid: util.ConvertToInt(a[1]), tyype: determineHandType(a[0])})
	}

	solution := 0
	slices.SortFunc(hands, compareCards)
	for i, v := range hands {
		solution += (i + 1) * v.bid
	}
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func compareCards(a, b hand) int {
	if a.tyype > b.tyype {
		return 1
	} else if a.tyype < b.tyype {
		return -1
	} else {
		aVals := changeValues(a.cards)
		bVals := changeValues(b.cards)

		for i := 0; i < len(a.cards); i++ {
			if aVals[i] > bVals[i] {
				return 1
			} else if aVals[i] < bVals[i] {
				return -1
			}
		}

		panic(fmt.Sprintf("What happens when two cards have the same everything? %v %v", aVals, bVals))
	}
}

func changeValues(cards string) []rune {
	res := make([]rune, 0)
	for _, v := range cards {
		switch v {
		case 'A':
			res = append(res, '9'+5)
		case 'K':
			res = append(res, '9'+4)
		case 'Q':
			res = append(res, '9'+3)
		case 'J':
			if part2 {
				res = append(res, '1')
			} else {
				res = append(res, '9'+2)
			}
		case 'T':
			res = append(res, '9'+1)
		default:
			res = append(res, v)
		}
	}

	return res
}

func determineHandType(cards string) handType {
	cnt := strings.Count(cards, "J")
	if part2 && cnt > 0 {
		if cnt == 5 {
			return fives
		}

		c := regexp.MustCompile("J+")
		cards = c.ReplaceAllString(cards, "")
	}
	cardMap := make(map[rune]int)
	for _, v := range cards {
		if _, ok := cardMap[v]; ok {
			cardMap[v]++
		} else {
			cardMap[v] = 1
		}
	}

	values := make([]int, 0)
	for _, v := range cardMap {
		values = append(values, v)
	}

	if part2 && cnt > 0 {
		max := slices.Max(values)
		values = make([]int, 0)
		add := false
		for _, v := range cardMap {
			if v == max && !add {
				v += cnt
				add = true
			}
			values = append(values, v)
		}
	}

	switch len(values) {
	case 1:
		return fives
	case 2:
		if slices.Contains(values, 4) {
			return fours
		} else {
			return fullhouse
		}
	case 3:
		if slices.Contains(values, 3) {
			return threes
		} else {
			return twopair
		}
	case 4:
		return pair
	default:
		return high
	}
}

type handType int

const (
	fives     handType = 6
	fours     handType = 5
	fullhouse handType = 4
	threes    handType = 3
	twopair   handType = 2
	pair      handType = 1
	high      handType = 0
)

type hand struct {
	cards string
	bid   int
	tyype handType
}
