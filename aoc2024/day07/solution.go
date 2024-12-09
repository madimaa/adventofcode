package day07

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/madimaa/adventofcode/aoc2024/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day07.txt")
	numbers := make(map[int][]int)
	for _, line := range input {
		target, _ := strconv.Atoi(strings.Split(line, ": ")[0])
		nums := make([]int, 0)
		for _, num := range strings.Split(strings.Split(line, ": ")[1], " ") {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}

		numbers[target] = nums
	}

	solution := 0
	for k, v := range numbers {
		perms := getPermutations(len(v) - 1)
		for _, perm := range perms {
			res := v[0]
			for i, r := range perm {
				switch r {
				case '+':
					res += v[i+1]
				case '×':
					res *= v[i+1]
				}
			}

			if res == k {
				solution += res
				break
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day07.txt")
	numbers := make(map[int][]int)
	for _, line := range input {
		target, _ := strconv.Atoi(strings.Split(line, ": ")[0])
		nums := make([]int, 0)
		for _, num := range strings.Split(strings.Split(line, ": ")[1], " ") {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}

		numbers[target] = nums
	}

	solution := 0
	for k, v := range numbers {
		perms := getPermutations2(len(v) - 1)
		for _, perm := range perms {
			res := v[0]
			for i, r := range perm {
				switch r {
				case '+':
					res += v[i+1]
				case '*':
					res *= v[i+1]
				case '|':
					res, _ = strconv.Atoi(fmt.Sprintf("%d%d", res, v[i+1]))
				}
			}

			if res == k {
				solution += res
				break
			}
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func getPermutations(length int) []string {
	if length == 1 {
		return []string{"+", "×"}
	}

	result := make([]string, 0)
	fill := getPermutations(length - 1)
	for _, v := range fill {
		result = append(result, fmt.Sprint("+", v))
		result = append(result, fmt.Sprint("×", v))
	}

	fill = nil
	return result
}

func getPermutations2(length int) []string {
	if length == 1 {
		return []string{"+", "*", "|"}
	}

	result := make([]string, 0)
	fill := getPermutations2(length - 1)
	for _, v := range fill {
		result = append(result, fmt.Sprint("+", v))
		result = append(result, fmt.Sprint("*", v))
		result = append(result, fmt.Sprint("|", v))
	}

	fill = nil
	return result
}
