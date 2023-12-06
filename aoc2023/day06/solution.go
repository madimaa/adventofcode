package day06

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day06.txt")
	trimmer := regexp.MustCompile("\\s+")
	races := make([]race, 0)
	times := strings.Split(trimmer.ReplaceAllString(input[0], " "), " ")
	distances := strings.Split(trimmer.ReplaceAllString(input[1], " "), " ")
	for i := 1; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			log.Fatal(err)
		}

		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			log.Fatal(err)
		}

		races = append(races, race{time: time, distance: distance})
	}

	solution := 1
	for _, r := range races {
		pass := 0
		for i := 0; i <= r.time; i++ {
			dist := i * (r.time - i)
			if dist > r.distance {
				pass++
			}
		}

		solution *= pass
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day06.txt")
	trimmer := regexp.MustCompile("\\s+")
	time, err := strconv.Atoi(strings.Split(trimmer.ReplaceAllString(input[0], ""), ":")[1])
	if err != nil {
		log.Fatal(err)
	}
	distance, err := strconv.Atoi(strings.Split(trimmer.ReplaceAllString(input[1], ""), ":")[1])
	if err != nil {
		log.Fatal(err)
	}

	solution := 0
	for i := 0; i <= time; i++ {
		dist := i * (time - i)
		if dist > distance {
			solution++
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

type race struct {
	time     int
	distance int
}
