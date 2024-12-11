package day25

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day25t.txt")
	nodes := make(map[string][]string)
	mKey := input[0][:3]
	for _, line := range input {
		key := line[:3]
		values := strings.Split(line[5:], " ")
		if _, ok := nodes[key]; !ok {
			nodes[key] = values
		} else {
			nodes[key] = append(nodes[key], values...)
		}

		for _, val := range values {
			if _, ok := nodes[val]; !ok {
				nodes[val] = make([]string, 0)
			}

			nodes[val] = append(nodes[val], key)
		}
	}

	keys := make([]string, 0, len(nodes))
	for k := range nodes {
		keys = append(keys, k)
	}

	for {
		keysCopy := make([]string, len(keys))
		copy(keysCopy, keys)
		rand.Shuffle(len(keysCopy), func(i, j int) { keysCopy[i], keysCopy[j] = keysCopy[j], keysCopy[i] })
		selectedKeys := keysCopy[:3]

		newNodes := make(map[string][]string, 0)
		for k, v := range nodes {
			newNodes[k] = v
		}
		for k, v := range newNodes {
			if slices.Contains(selectedKeys, k) {
				vTemp := make([]string, len(v))
				copy(vTemp, v)
				rand.Shuffle(len(vTemp), func(i, j int) { vTemp[i], vTemp[j] = vTemp[j], vTemp[i] })
				for i, node := range vTemp {
					if !slices.Contains(selectedKeys, node) {
						vTemp = slices.Delete(vTemp, i, i+1)
						newNodes[k] = vTemp

						temp := make([]string, len(newNodes[node]))
						copy(temp, newNodes[node])

						break
					}
				}
			}
		}

		if !checkReachEveryNode(newNodes, mKey) {
			fmt.Println(keysCopy[:3])
			break
		}
	}

	solution := 0
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	//input := util.OpenFile("assets/day25.txt")

	solution := 0
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func checkReachEveryNode(nodes map[string][]string, mKey string) bool {
	list := make([]string, 0)
	list = append(list, mKey)
	index := 0
	for {
		if index >= len(list) {
			break
		}
		node := list[index]
		for _, v := range nodes[node] {
			if !slices.Contains(list, v) {
				list = append(list, v)
			}
		}

		index++
	}

	return len(list) == len(nodes)
}
