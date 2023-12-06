package day05

import (
	"log"
	"math"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day05.txt")
	alma := almanac{seeds: make([]int64, 0), seedToSoil: make([]amap, 0), soilToFertilizer: make([]amap, 0), fertilizerToWater: make([]amap, 0), waterToLight: make([]amap, 0),
		lightToTemperature: make([]amap, 0), temperatureToHumidity: make([]amap, 0), humidityToLocation: make([]amap, 0)}
	seeds := strings.Split(input[0], " ")
	for i := 1; i < len(seeds); i++ {
		alma.seeds = append(alma.seeds, util.ConvertToInt64(seeds[i]))
	}

	last := ""
	for i := 2; i < len(input); i++ {
		if len(input[i]) == 0 {
			continue
		}

		line := strings.Split(input[i], " ")
		if strings.HasSuffix(input[i], ":") {
			last = strings.Split(line[0], "-")[0]
			continue
		}
		a := amap{destination: util.ConvertToInt64(line[0]), source: util.ConvertToInt64(line[1]), rangez: util.ConvertToInt64(line[2])}
		switch last {
		case "seed":
			alma.seedToSoil = append(alma.seedToSoil, a)
		case "soil":
			alma.soilToFertilizer = append(alma.soilToFertilizer, a)
		case "fertilizer":
			alma.fertilizerToWater = append(alma.fertilizerToWater, a)
		case "water":
			alma.waterToLight = append(alma.waterToLight, a)
		case "light":
			alma.lightToTemperature = append(alma.lightToTemperature, a)
		case "temperature":
			alma.temperatureToHumidity = append(alma.temperatureToHumidity, a)
		case "humidity":
			alma.humidityToLocation = append(alma.humidityToLocation, a)
		}
	}
	solution := int64(math.MaxInt64)
	for _, seed := range alma.seeds {
		soil := searchMap(alma.seedToSoil, seed)
		fertilizer := searchMap(alma.soilToFertilizer, soil)
		water := searchMap(alma.fertilizerToWater, fertilizer)
		light := searchMap(alma.waterToLight, water)
		temperature := searchMap(alma.lightToTemperature, light)
		humidity := searchMap(alma.temperatureToHumidity, temperature)
		location := searchMap(alma.humidityToLocation, humidity)

		if solution > location {
			solution = location
		}
	}

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day05.txt")
	alma := almanac{seeds: make([]int64, 0), seedToSoil: make([]amap, 0), soilToFertilizer: make([]amap, 0), fertilizerToWater: make([]amap, 0), waterToLight: make([]amap, 0),
		lightToTemperature: make([]amap, 0), temperatureToHumidity: make([]amap, 0), humidityToLocation: make([]amap, 0)}
	seeds := strings.Split(input[0], " ")
	for i := 1; i < len(seeds); i++ {
		alma.seeds = append(alma.seeds, util.ConvertToInt64(seeds[i]))
	}

	last := ""
	for i := 2; i < len(input); i++ {
		if len(input[i]) == 0 {
			continue
		}

		line := strings.Split(input[i], " ")
		if strings.HasSuffix(input[i], ":") {
			last = strings.Split(line[0], "-")[0]
			continue
		}
		a := amap{destination: util.ConvertToInt64(line[0]), source: util.ConvertToInt64(line[1]), rangez: util.ConvertToInt64(line[2])}
		switch last {
		case "seed":
			alma.seedToSoil = append(alma.seedToSoil, a)
		case "soil":
			alma.soilToFertilizer = append(alma.soilToFertilizer, a)
		case "fertilizer":
			alma.fertilizerToWater = append(alma.fertilizerToWater, a)
		case "water":
			alma.waterToLight = append(alma.waterToLight, a)
		case "light":
			alma.lightToTemperature = append(alma.lightToTemperature, a)
		case "temperature":
			alma.temperatureToHumidity = append(alma.temperatureToHumidity, a)
		case "humidity":
			alma.humidityToLocation = append(alma.humidityToLocation, a)
		}
	}

	var solution int64
	for location := int64(0); location < int64(math.MaxInt64); location++ {
		humidity := searchMapReverse(alma.humidityToLocation, location)
		temperature := searchMapReverse(alma.temperatureToHumidity, humidity)
		light := searchMapReverse(alma.lightToTemperature, temperature)
		water := searchMapReverse(alma.waterToLight, light)
		fertilizer := searchMapReverse(alma.fertilizerToWater, water)
		soil := searchMapReverse(alma.soilToFertilizer, fertilizer)
		seed := searchMapReverse(alma.seedToSoil, soil)
		//log.Print(seed, soil, fertilizer, water, light, temperature, humidity, location)
		for i := 0; i < len(alma.seeds); i += 2 {
			init := alma.seeds[i]
			rangez := alma.seeds[i+1]
			if seed >= init && seed <= init+rangez {
				solution = location
				goto FINISH
			}
		}
	}

FINISH:
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func searchMap(aToBMap []amap, input int64) int64 {
	for _, m := range aToBMap {
		if m.source <= input && m.rangez+m.source > input {
			return input - m.source + m.destination
		}
	}

	return input
}

func searchMapReverse(aToBMap []amap, output int64) int64 {
	for _, m := range aToBMap {
		if m.destination <= output && m.rangez+m.destination > output {
			return output - m.destination + m.source
		}
	}

	return output
}

type almanac struct {
	seeds []int64

	seedToSoil            []amap
	soilToFertilizer      []amap
	fertilizerToWater     []amap
	waterToLight          []amap
	lightToTemperature    []amap
	temperatureToHumidity []amap
	humidityToLocation    []amap
}

type amap struct {
	destination int64
	source      int64
	rangez      int64
}
