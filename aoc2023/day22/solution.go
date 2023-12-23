package day22

import (
	"cmp"
	"log"
	"regexp"
	"slices"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day22.txt")
	c := regexp.MustCompile(`([0-9]+),([0-9]+),([0-9]+)~([0-9]+),([0-9]+),([0-9]+)`)
	bricks := make([]brick, 0)
	for _, line := range input {
		groups := c.FindStringSubmatch(line)
		b := brick{util.ConvertToInt(groups[1]), util.ConvertToInt(groups[2]), util.ConvertToInt(groups[3]),
			util.ConvertToInt(groups[4]), util.ConvertToInt(groups[5]), util.ConvertToInt(groups[6])}
		bricks = append(bricks, b)
	}

	slices.SortFunc(bricks, func(a, b brick) int { return cmp.Compare(a.z1, b.z1) })
	for i, b := range bricks {
		fall(b, bricks, i)
	}

	soloSupport := make([]brick, 0)
	for i := len(bricks) - 1; i >= 0; i-- {
		supported := make([]brick, 0)
		b := bricks[i]
		for j := i - 1; j >= 0; j-- {
			if b.supportedBy(bricks[j]) {
				supported = append(supported, bricks[j])
			}
		}

		if len(supported) == 1 {
			if !slices.Contains(soloSupport, supported[0]) {
				soloSupport = append(soloSupport, supported[0])
			}
		}
	}

	solution := len(bricks) - len(soloSupport)
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day22.txt")
	c := regexp.MustCompile(`([0-9]+),([0-9]+),([0-9]+)~([0-9]+),([0-9]+),([0-9]+)`)
	bricks := make([]brick, 0)
	for _, line := range input {
		groups := c.FindStringSubmatch(line)
		b := brick{util.ConvertToInt(groups[1]), util.ConvertToInt(groups[2]), util.ConvertToInt(groups[3]),
			util.ConvertToInt(groups[4]), util.ConvertToInt(groups[5]), util.ConvertToInt(groups[6])}
		bricks = append(bricks, b)
	}

	slices.SortFunc(bricks, func(a, b brick) int { return cmp.Compare(a.z1, b.z1) })
	for i, b := range bricks {
		fall(b, bricks, i)
	}

	soloSupport := make([]brick, 0)
	for i := len(bricks) - 1; i >= 0; i-- {
		supported := make([]brick, 0)
		b := bricks[i]
		for j := i - 1; j >= 0; j-- {
			if b.supportedBy(bricks[j]) {
				supported = append(supported, bricks[j])
			}
		}

		if len(supported) == 1 {
			if !slices.Contains(soloSupport, supported[0]) {
				soloSupport = append(soloSupport, supported[0])
			}
		}
	}

	solution := 0
	for _, v := range soloSupport {
		cbricks := slices.Clone(bricks)
		vIndex := slices.Index(cbricks, v)
		cbricks = slices.Delete(cbricks, vIndex, vIndex+1)

		for i, b := range cbricks {
			nb := b
			fall(b, cbricks, i)
			if nb.z1 != cbricks[i].z1 {
				solution++
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func fall(b brick, bricks []brick, index int) {
	for b.z1 > 1 {
		for i := index - 1; i >= 0; i-- {
			b2 := bricks[i]
			if b.supportedBy(b2) {
				return
			}
		}
		b.z1--
		b.z2--
		bricks[index] = b
	}
}

type brick struct {
	x1, y1, z1 int
	x2, y2, z2 int
}

func (b1 brick) supportedBy(b2 brick) bool {
	if b1.z1-b2.z2 == 1 {
		a := b2.x1 <= b1.x2 && b2.x2 >= b1.x1
		b := b2.y1 <= b1.y2 && b2.y2 >= b1.y1
		if a && b {
			return true
		}
	}

	return false
}
