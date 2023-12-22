package day18

import (
	"image"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day18.txt")
	vertices := make([]image.Point, 0)
	vertices = append(vertices, image.Point{0, 0})
	picks := 0
	for _, line := range input {
		fields := strings.Fields(line)
		last := vertices[len(vertices)-1]
		value := util.ConvertToInt(fields[1])
		picks += value
		switch fields[0] {
		case "U":
			vertices = append(vertices, image.Point{last.X, last.Y - value})
		case "R":
			vertices = append(vertices, image.Point{last.X + value, last.Y})
		case "D":
			vertices = append(vertices, image.Point{last.X, last.Y + value})
		case "L":
			vertices = append(vertices, image.Point{last.X - value, last.Y})
		}
	}

	solution := shoelaceFormula(vertices)
	log.SetFlags(0)
	log.Printf("%.f", solution+float64(picks/2)+1.0)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day18.txt")
	vertices := make([]image.Point, 0)
	vertices = append(vertices, image.Point{0, 0})
	picks := 0
	for _, line := range input {
		fields := strings.Fields(line)
		hex := fields[2][1:8]
		last := vertices[len(vertices)-1]
		v, err := strconv.ParseInt(hex[1:6], 16, 0)
		util.PanicOnError(err)
		value := int(v)
		picks += value
		switch hex[len(hex)-1] {
		case '3':
			vertices = append(vertices, image.Point{last.X, last.Y - value})
		case '0':
			vertices = append(vertices, image.Point{last.X + value, last.Y})
		case '1':
			vertices = append(vertices, image.Point{last.X, last.Y + value})
		case '2':
			vertices = append(vertices, image.Point{last.X - value, last.Y})
		}
	}

	solution := shoelaceFormula(vertices)
	log.SetFlags(0)
	log.Printf("%.f", solution+float64(picks/2)+1.0)
}

func shoelaceFormula(vertices []image.Point) float64 {
	area := 0

	j := len(vertices) - 1
	for i := 0; i < len(vertices); i++ {
		area += (vertices[j].X + vertices[i].X) * (vertices[j].Y - vertices[i].Y)

		j = i
	}

	return math.Abs(float64(area) / 2.0)
}
