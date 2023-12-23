package day19

import (
	"log"
	"maps"
	"regexp"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day19.txt")
	c := regexp.MustCompile(`([a-z]+)\{(.*)\}`)
	workflows := make(map[string][]condition)
	i := 0
	for ; i < len(input); i++ {
		if len(input[i]) == 0 {
			i++
			break
		}

		groups := c.FindStringSubmatch(input[i])
		workflows[groups[1]] = make([]condition, 0)
		for _, v := range strings.Split(groups[2], ",") {
			if strings.Contains(v, ":") {
				a := strings.Split(v, ":")
				b := a[0]
				r := []rune(b)[0]
				val := util.ConvertToInt(b[2:])
				f := func(in int) bool {
					switch b[1] {
					case '>':
						if in > val {
							return true
						} else {
							return false
						}
					case '<':
						if in < val {
							return true
						} else {
							return false
						}
					}

					panic("Something went wrong!")
				}

				c := condition{vType: r, value: val, check: f, goal: a[1]}
				workflows[groups[1]] = append(workflows[groups[1]], c)
			} else {
				workflows[groups[1]] = append(workflows[groups[1]], condition{vType: '0', check: func(in int) bool { return true }, goal: v})
			}
		}
	}

	solution := 0
	c2 := regexp.MustCompile(`\{x=([0-9]+),m=([0-9]+),a=([0-9]+),s=([0-9]+)\}`)
	for ; i < len(input); i++ {
		groups := c2.FindStringSubmatch(input[i])
		p := map[rune]int{'x': util.ConvertToInt(groups[1]), 'm': util.ConvertToInt(groups[2]), 'a': util.ConvertToInt(groups[3]), 's': util.ConvertToInt(groups[4])}
		solution += process(workflows, "in", p)
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day19.txt")
	c := regexp.MustCompile(`([a-z]+)\{(.*)\}`)
	workflows := make(map[string][]condition)
	i := 0
	for ; i < len(input); i++ {
		if len(input[i]) == 0 {
			i++
			break
		}

		groups := c.FindStringSubmatch(input[i])
		workflows[groups[1]] = make([]condition, 0)
		for _, v := range strings.Split(groups[2], ",") {
			if strings.Contains(v, ":") {
				a := strings.Split(v, ":")
				b := a[0]
				r := []rune(b)[0]
				val := util.ConvertToInt(b[2:])
				f := func(in int) bool {
					switch b[1] {
					case '>':
						if in > val {
							return true
						} else {
							return false
						}
					case '<':
						if in < val {
							return true
						} else {
							return false
						}
					}

					panic("Something went wrong!")
				}

				c := condition{vType: r, value: val, check: f, goal: a[1], relation: []rune(b)[1]}
				workflows[groups[1]] = append(workflows[groups[1]], c)
			} else {
				workflows[groups[1]] = append(workflows[groups[1]], condition{vType: '0', check: func(in int) bool { return true }, goal: v})
			}
		}
	}

	ranges := map[rune][2]int{'x': {1, 4000}, 'm': {1, 4000}, 'a': {1, 4000}, 's': {1, 4000}}

	solution := calcRanges(workflows, "in", ranges)
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func process(workflows map[string][]condition, key string, p map[rune]int) int {
	conditions := workflows[key]
	for _, c := range conditions {
		if c.check(p[c.vType]) {
			if c.goal == "R" {
				return 0
			} else if c.goal == "A" {
				return p['x'] + p['m'] + p['a'] + p['s']
			} else {
				return process(workflows, c.goal, p)
			}
		}
	}

	panic("Something went wrong!")
}

func calcRanges(workflows map[string][]condition, key string, ranges map[rune][2]int) int {
	result := 0
	if key == "R" {
		return 0
	}
	if key == "A" {
		res := 1
		for _, v := range ranges {
			res *= v[1] - v[0] + 1
		}
		return res
	}

	for _, c := range workflows[key] {
		if c.vType == '0' {
			if c.goal == "R" {
				return result
			} else if c.goal == "A" || c.goal == "R" {
				res := 1
				for _, v := range ranges {
					res *= v[1] - v[0] + 1
				}
				result += res
			} else {
				result += calcRanges(workflows, c.goal, ranges)
			}
		} else {
			newRanges := maps.Clone(ranges)
			if c.relation == '<' {
				newRanges[c.vType] = [2]int{newRanges[c.vType][0], min(newRanges[c.vType][1], c.value-1)}
				ranges[c.vType] = [2]int{max(newRanges[c.vType][0], c.value), ranges[c.vType][1]}
			} else {
				newRanges[c.vType] = [2]int{max(newRanges[c.vType][0], c.value+1), newRanges[c.vType][1]}
				ranges[c.vType] = [2]int{ranges[c.vType][0], min(newRanges[c.vType][1], c.value)}
			}

			result += calcRanges(workflows, c.goal, newRanges)
		}
	}

	return result
}

type condition struct {
	vType    rune
	value    int
	check    func(v int) bool
	goal     string
	relation rune
}
