package day20

import (
	"log"
	"strings"

	"github.com/madimaa/adventofcode/aoc2023/util"
)

func Part1() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day20.txt")
	modules := make(map[string]module)
	for _, line := range input {
		key := strings.Split(line, " -> ")[0]
		values := strings.Split(strings.Split(line, " -> ")[1], ", ")
		if key == "broadcaster" {
			modules[key] = module{typez: bc, connections: values}
		} else {
			switch []rune(key)[0] {
			case '%':
				modules[key[1:]] = module{typez: flipflop, state: 0, connections: values}
			case '&':
				modules[key[1:]] = module{typez: conjunction, state: 0, inputs: make(map[string]signal), connections: values}
			}
		}
	}

	for key, mod := range modules {
		for _, v := range mod.connections {
			if modules[v].typez == conjunction {
				if _, ok := modules[v].inputs[key]; !ok {
					modules[v].inputs[key] = 0
				}
			}
		}
	}

	lows, highs := 0, 0
	for i := 0; i < 1000; i++ {
		l, h := process(modules, "broadcaster", low)
		lows += l
		highs += h
	}

	solution := lows * highs
	log.SetFlags(0)
	log.Printf("%d", solution)
}

func Part2() {
	log.SetFlags(log.Lshortfile)
	input := util.OpenFile("assets/day20.txt")
	modules := make(map[string]module)
	for _, line := range input {
		key := strings.Split(line, " -> ")[0]
		values := strings.Split(strings.Split(line, " -> ")[1], ", ")
		if key == "broadcaster" {
			modules[key] = module{typez: bc, connections: values}
		} else {
			switch []rune(key)[0] {
			case '%':
				modules[key[1:]] = module{typez: flipflop, state: 0, connections: values}
			case '&':
				modules[key[1:]] = module{typez: conjunction, state: 0, inputs: make(map[string]signal), connections: values}
			}
		}
	}

	rxSource := ""
	for key, mod := range modules {
		for _, v := range mod.connections {
			if modules[v].typez == conjunction {
				if _, ok := modules[v].inputs[key]; !ok {
					modules[v].inputs[key] = 0
				}
			}

			if v == "rx" {
				rxSource = key
			}
		}
	}

	rxConns := make(map[string]int)
	for k := range modules[rxSource].inputs {
		rxConns[k] = 0
	}
	solution := 1
	for {
		key := process2(modules, "broadcaster", low, rxConns)
		if len(key) > 0 {
			rxConns[key] = solution
		}

		shouldBreak := true
		for _, v := range rxConns {
			if v == 0 {
				shouldBreak = false
			}
		}

		if shouldBreak {
			break
		}

		solution++
	}

	vals := make([]int, 0)
	for _, v := range rxConns {
		vals = append(vals, v)
	}

	solution = util.LCM(vals[0], vals[1], vals[2:]...)

	log.SetFlags(0)
	log.Printf("%d", solution)
}

func process(modules map[string]module, key string, sig signal) (highs, lows int) {
	lows++
	sources := make([]string, 0)
	outputs := make([]string, 0)
	signals := make([]signal, 0)
	for _, v := range modules["broadcaster"].connections {
		sources = append(sources, "boradcaster")
		outputs = append(outputs, v)
		signals = append(signals, 0)
	}

	for len(outputs) > 0 {
		source := sources[0]
		sources = sources[1:]
		key := outputs[0]
		outputs = outputs[1:]
		sig := signals[0]
		signals = signals[1:]

		if sig == low {
			lows++
		} else {
			highs++
		}

		if _, ok := modules[key]; !ok {
			continue
		}

		mod := modules[key]

		switch mod.typez {
		case flipflop:
			if sig == low {
				if mod.state == 0 {
					mod.state = 1
					for _, v := range mod.connections {
						signals = append(signals, high)
						outputs = append(outputs, v)
						sources = append(sources, key)
					}
				} else {
					mod.state = 0
					for _, v := range mod.connections {
						signals = append(signals, low)
						outputs = append(outputs, v)
						sources = append(sources, key)
					}
				}

				modules[key] = mod
			}
		case conjunction:
			mod.inputs[source] = sig
			onlyHigh := true
			for _, v := range mod.inputs {
				if v != high {
					onlyHigh = false
					break
				}
			}

			if onlyHigh {
				for _, v := range mod.connections {
					signals = append(signals, low)
					outputs = append(outputs, v)
					sources = append(sources, key)
				}
			} else {
				for _, v := range mod.connections {
					signals = append(signals, high)
					outputs = append(outputs, v)
					sources = append(sources, key)
				}
			}
		}
	}

	return
}

func process2(modules map[string]module, key string, sig signal, rxConns map[string]int) string {
	sources := make([]string, 0)
	outputs := make([]string, 0)
	signals := make([]signal, 0)
	for _, v := range modules["broadcaster"].connections {
		sources = append(sources, "boradcaster")
		outputs = append(outputs, v)
		signals = append(signals, 0)
	}

	for len(outputs) > 0 {
		source := sources[0]
		sources = sources[1:]
		key := outputs[0]
		outputs = outputs[1:]
		sig := signals[0]
		signals = signals[1:]

		if v, ok := rxConns[source]; ok && key == "gf" && v == 0 && sig == high {
			return source
		}
		if _, ok := modules[key]; !ok {
			continue
		}

		mod := modules[key]

		switch mod.typez {
		case flipflop:
			if sig == low {
				if mod.state == 0 {
					mod.state = 1
					for _, v := range mod.connections {
						signals = append(signals, high)
						outputs = append(outputs, v)
						sources = append(sources, key)
					}
				} else {
					mod.state = 0
					for _, v := range mod.connections {
						signals = append(signals, low)
						outputs = append(outputs, v)
						sources = append(sources, key)
					}
				}

				modules[key] = mod
			}
		case conjunction:
			mod.inputs[source] = sig
			onlyHigh := true
			for _, v := range mod.inputs {
				if v != high {
					onlyHigh = false
					break
				}
			}

			if onlyHigh {
				for _, v := range mod.connections {
					signals = append(signals, low)
					outputs = append(outputs, v)
					sources = append(sources, key)
				}
			} else {
				for _, v := range mod.connections {
					signals = append(signals, high)
					outputs = append(outputs, v)
					sources = append(sources, key)
				}
			}
		}
	}

	return ""
}

type signal int
type modType int

const (
	bc          modType = -1
	flipflop    modType = 0
	conjunction modType = 1

	low  signal = 0
	high signal = 1
)

type module struct {
	typez       modType
	state       int
	inputs      map[string]signal
	connections []string
}
