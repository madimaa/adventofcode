package day02

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day02.txt")
	trimmer := regexp.MustCompile("\\s+")
	games := make([]game, 0)
	for _, line := range input {
		line = trimmer.ReplaceAllString(line, " ")
		a := strings.Split(line, ": ")
		gameId, err := strconv.Atoi(strings.Split(a[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		g := game{id: gameId, draws: make([]color, 0)}
		for _, draw := range strings.Split(a[1], "; ") {
			for _, v := range strings.Split(draw, ", ") {
				b := strings.Split(v, " ")
				cnt, err := strconv.Atoi(b[0])
				if err != nil {
					log.Fatal(err)
				}

				c := color{}
				switch b[1] {
				case "red":
					c.red = cnt
				case "green":
					c.green = cnt
				case "blue":
					c.blue = cnt
				}

				g.draws = append(g.draws, c)
			}
		}
		games = append(games, g)
	}

	sum := 0
	for _, v := range games {
		possible := true
		for _, set := range v.draws {
			if set.blue > 14 || set.green > 13 || set.red > 12 {
				possible = false
			}
		}

		if possible {
			sum += v.id
		}
	}

	log.SetFlags(0)
	log.Printf("%d", sum)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day02.txt")
	trimmer := regexp.MustCompile("\\s+")
	games := make([]game, 0)
	for _, line := range input {
		line = trimmer.ReplaceAllString(line, " ")
		a := strings.Split(line, ": ")
		gameId, err := strconv.Atoi(strings.Split(a[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		g := game{id: gameId, draws: make([]color, 0)}
		for _, draw := range strings.Split(a[1], "; ") {
			for _, v := range strings.Split(draw, ", ") {
				b := strings.Split(v, " ")
				cnt, err := strconv.Atoi(b[0])
				if err != nil {
					log.Fatal(err)
				}

				c := color{}
				switch b[1] {
				case "red":
					c.red = cnt
				case "green":
					c.green = cnt
				case "blue":
					c.blue = cnt
				}

				g.draws = append(g.draws, c)
			}
		}
		games = append(games, g)
	}

	sum := 0
	for _, v := range games {
		maxGreen := 0
		maxBlue := 0
		maxRed := 0
		for _, set := range v.draws {
			if maxGreen < set.green {
				maxGreen = set.green
			}
			if maxBlue < set.blue {
				maxBlue = set.blue
			}
			if maxRed < set.red {
				maxRed = set.red
			}
		}

		sum += maxGreen * maxBlue * maxRed
	}

	log.SetFlags(0)
	log.Printf("%d", sum)
}

type game struct {
	id    int
	draws []color
}

type color struct {
	red   int
	green int
	blue  int
}
