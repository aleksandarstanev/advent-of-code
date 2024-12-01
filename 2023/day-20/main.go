package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Low = iota
	High
)

type Module interface {
	GetPulses(pulse Pulse) []Pulse
	GetDestinations() []string
}

type Pulse struct {
	pulseType int
	module    string
}

type Broadcaster struct {
	destinations []string
}

func (b *Broadcaster) GetPulses(pulse Pulse) []Pulse {
	pulses := make([]Pulse, 0)

	for _, dest := range b.destinations {
		pulses = append(pulses, Pulse{
			pulse.pulseType,
			dest,
		})
	}

	return pulses
}

func (b Broadcaster) GetDestinations() []string {
	return b.destinations
}

type FlipFlop struct {
	state        bool
	destinations []string
}

func (f *FlipFlop) GetPulses(pulse Pulse) []Pulse {
	pulses := make([]Pulse, 0)

	if pulse.pulseType == High {
		return pulses
	}

	if f.state {
		for _, dest := range f.destinations {
			pulses = append(pulses, Pulse{
				Low,
				dest,
			})
		}
	} else {
		for _, dest := range f.destinations {
			pulses = append(pulses, Pulse{
				High,
				dest,
			})
		}
	}

	f.state = !f.state

	return pulses
}

func (f FlipFlop) GetDestinations() []string {
	return f.destinations
}

type Conjunction struct {
	inputPulses  map[string]int
	destinations []string
}

func (c *Conjunction) GetPulses(pulse Pulse) []Pulse {
	pulses := make([]Pulse, 0)

	// if pulse.module == "hb" {
	// 	ones := 0
	// 	for _, v := range c.inputPulses {
	// 		if v == High {
	// 			ones += 1
	// 		}
	// 	}

	// 	if ones > 1 {
	// 		fmt.Println(c.inputPulses)
	// 	}
	// }

	allHigh := true
	for _, val := range c.inputPulses {
		if val == Low {
			allHigh = false
		}
	}

	if allHigh {
		for _, dest := range c.destinations {
			pulses = append(pulses, Pulse{
				Low,
				dest,
			})
		}
	} else {
		for _, dest := range c.destinations {
			pulses = append(pulses, Pulse{
				High,
				dest,
			})
		}
	}

	return pulses
}

func (c Conjunction) GetDestinations() []string {
	return c.destinations
}

func solvePart1(nameToModule map[string]Module) int {
	const STEPS = 1000

	low, high := 0, 0

	for i := 0; i < STEPS; i++ {
		queue := make([]Pulse, 0)

		low += 1
		queue = append(queue, Pulse{
			Low,
			"broadcaster",
		})

		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]

			// fmt.Println(pulse)

			module, ok := nameToModule[pulse.module]
			if !ok {
				continue
			}

			newPulses := module.GetPulses(pulse)

			for _, newPulse := range newPulses {
				if newPulse.pulseType == Low {
					low += 1
				} else {
					high += 1
				}

				queue = append(queue, newPulse)

				value, ok := nameToModule[newPulse.module].(*Conjunction)
				if ok {
					value.inputPulses[pulse.module] = newPulse.pulseType
				}
			}
		}
	}

	return low * high
}

func solvePart2(nameToModule map[string]Module) int {
	// 11752068000 is too low
	// 11752068001 is too low

	steps := 0
	isRxFound := false
	for !isRxFound {
		queue := make([]Pulse, 0)

		queue = append(queue, Pulse{
			Low,
			"broadcaster",
		})

		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]

			if pulse.module == "rx" && pulse.pulseType == Low {
				isRxFound = true
				break
			}

			module, ok := nameToModule[pulse.module]
			if !ok {
				continue
			}

			newPulses := module.GetPulses(pulse)

			for _, newPulse := range newPulses {
				queue = append(queue, newPulse)

				value, ok := nameToModule[newPulse.module].(*Conjunction)
				if ok {
					value.inputPulses[pulse.module] = newPulse.pulseType

					if newPulse.module == "hb" {
						// ones := 0
						// for _, v := range value.inputPulses {
						// 	if v == High {
						// 		ones += 1
						// 	}
						// }

						if value.inputPulses["rr"] == High {
							fmt.Println(value.inputPulses, steps)
						}
					}
				}
			}
		}

		steps += 1
	}

	return steps
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nameToModule := make(map[string]Module)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " -> ")

		name := parts[0]
		destinations := strings.Split(parts[1], ", ")
		if name == "broadcaster" {
			nameToModule[name] = &Broadcaster{
				destinations,
			}
		} else if name[0] == '%' {
			nameToModule[name[1:]] = &FlipFlop{
				false,
				destinations,
			}
		} else if name[0] == '&' {
			nameToModule[name[1:]] = &Conjunction{
				make(map[string]int), // to fill
				destinations,
			}
		}
	}

	for key, value := range nameToModule {
		destinations := value.GetDestinations()

		for _, dest := range destinations {
			value, ok := nameToModule[dest].(*Conjunction)
			if ok {
				value.inputPulses[key] = Low
			}
		}
	}

	ans := solvePart2(nameToModule)
	fmt.Println("Answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
