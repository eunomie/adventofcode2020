package main

import (
	"testing"

	"github.com/eunomie/adventofcode2020/lib"
)

func TestDockingValues(t *testing.T) {
	const input = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

	c := NewDockingComputer()
	c.Run(lib.AsStringArray(input))
	v := c.SumValuesInMemory()
	if v != 165 {
		t.Fatal(v, 165)
	}
}
