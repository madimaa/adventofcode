package day08

import (
	"log"
	"regexp"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day08.txt")
	instructions := input[0]
	c := regexp.MustCompile(`([A-Z]+)\s=\s\(([A-Z]+),\s([A-Z]+)\)`)
	directions := make(map[string]_map)
	for i := 2; i < len(input); i++ {
		groups := c.FindStringSubmatch(input[i])
		directions[groups[1]] = _map{left: groups[2], right: groups[3]}
	}

	node := "AAA"
	solution := 0
	for {
		if node == "ZZZ" {
			break
		}

		instruction := []rune(instructions)[solution%len(instructions)]
		switch instruction {
		case 'L':
			node = directions[node].left
		case 'R':
			node = directions[node].right
		}
		solution++
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day08.txt")
	instructions := input[0]
	c := regexp.MustCompile(`([A-Z]+)\s=\s\(([A-Z]+),\s([A-Z]+)\)`)
	directions := make(map[string]_map)
	nodes := make([]string, 0)
	for i := 2; i < len(input); i++ {
		groups := c.FindStringSubmatch(input[i])
		node := groups[1]
		directions[node] = _map{left: groups[2], right: groups[3]}
		if strings.HasSuffix(node, "A") {
			nodes = append(nodes, node)
		}
	}

	stepList := make([]int, 0)
	for i := range nodes {
		steps := 0
		for {
			if strings.HasSuffix(nodes[i], "Z") {
				break
			}

			instruction := []rune(instructions)[steps%len(instructions)]
			switch instruction {
			case 'L':
				nodes[i] = directions[nodes[i]].left
			case 'R':
				nodes[i] = directions[nodes[i]].right
			}

			steps++
		}

		stepList = append(stepList, steps)
	}

	solution := util.LCM(stepList[0], stepList[1], stepList[2:]...)
	log.SetFlags(0)
	log.Printf("%d", solution)
}

type _map struct {
	left  string
	right string
}
