package main

import (
	"fmt"
	"math"
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
	memory map[int64]int
}

func NewDockingComputer() *DockingComputer {
	return &DockingComputer{
		memory: map[int64]int{},
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
		value, _ := strconv.Atoi(match[2])
		address, _ := strconv.Atoi(match[1])
		binValue := fmt.Sprintf("%036b", address)
		runes := []rune(binValue)
		vars := []int{}
		for i, m := range mask {
			if m == '1' {
				runes[i] = '1'
			} else if m == 'X' {
				vars = append(vars, i)
			}
		}
		combinations(runes, vars, func(address int64) {
			c.memory[address] = value
		})
	}
}

func combinations(runes []rune, vars []int, callback func(address int64)) {
	nbCombinations := int(math.Pow(2., float64(len(vars))))

	combPattern := fmt.Sprintf("%%0%db", len(vars))

	for i := 0; i < nbCombinations; i++ {
		combStr := fmt.Sprintf(combPattern, i)
		comb := []rune(combStr)
		combIdx := 0
		for _, idx := range vars {
			runes[idx] = comb[combIdx]
			combIdx++
		}
		binValue := string(runes)
		intAddress, _ := strconv.ParseInt(binValue, 2, 64)

		callback(intAddress)
	}
}

func (c *DockingComputer) SumValuesInMemory() int {
	var sum int
	for _, v := range c.memory {
		sum += v
	}
	return sum
}
