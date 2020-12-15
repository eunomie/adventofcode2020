package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	c := NewDockingComputer()
	c.Run(lib.AsStringArray(lib.Input()))
	v := c.SumValuesInMemory()
	fmt.Println("sum:", v)
}

type DockingComputer struct {
	memory map[int]int64
}

func NewDockingComputer() *DockingComputer {
	return &DockingComputer{
		memory: map[int]int64{},
	}
}

func (c *DockingComputer) Run(instructions []string) {
	pattern := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	mask := ""
	for _, instruction := range instructions {
		if strings.HasPrefix(instruction, "mask = ") {
			mask = strings.TrimPrefix(instruction, "mask = ")
			continue
		}
		match := pattern.FindStringSubmatch(instruction)
		address, _ := strconv.Atoi(match[1])
		value, _ := strconv.Atoi(match[2])
		binValue := fmt.Sprintf("%036b", value)
		runes := []rune(binValue)
		for i, m := range mask {
			if m == '0' {
				runes[i] = '0'
			} else if m == '1' {
				runes[i] = '1'
			}
		}
		binValue = string(runes)

		intVal, _ := strconv.ParseInt(binValue, 2, 64)
		c.memory[address] = intVal
	}
}

func (c *DockingComputer) SumValuesInMemory() int64 {
	var sum int64
	for _, v := range c.memory {
		sum += v
	}
	return sum
}
