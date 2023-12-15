package day15

import (
	"log"
	"slices"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day15.txt")
	solution := 0
	for _, v := range strings.Split(input[0], ",") {
		solution += hash(v)
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day15.txt")
	boxes := make([][]lens, 256)
	for i := range boxes {
		boxes[i] = make([]lens, 0)
	}
	for _, v := range strings.Split(input[0], ",") {
		if strings.HasSuffix(v, "-") {
			label := strings.TrimSuffix(v, "-")
			location := hash(label)
			index := -1
			for i, b := range boxes[location] {
				if b.label == label {
					index = i
					break
				}
			}

			if index != -1 {
				boxes[location] = slices.Delete(boxes[location], index, index+1)
			}
		} else {
			a := strings.Split(v, "=")
			label := a[0]
			fl := util.ConvertToInt(a[1])
			location := hash(label)

			index := -1
			for i, b := range boxes[location] {
				if b.label == label {
					index = i
					break
				}
			}

			l := lens{label, fl}
			if index == -1 {
				boxes[location] = append(boxes[location], l)
			} else {
				boxes[location][index] = l
			}
		}
	}

	solution := 0
	for index, b := range boxes {
		for slot, l := range b {
			log.Printf("Box: %d, slot: %d, focal length: %d", (index + 1), (slot + 1), l.focal)
			solution += (index + 1) * (slot + 1) * l.focal
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func hash(input string) int {
	current := 0
	for _, v := range input {
		current += int(v)
		current *= 17
		current %= 256
	}

	return current
}

type lens struct {
	label string
	focal int
}
