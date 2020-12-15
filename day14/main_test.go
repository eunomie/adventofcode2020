package main

import (
	"testing"

	"github.com/eunomie/adventofcode2020/lib"
)

func TestDockingValues(t *testing.T) {
	const input = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

	c := NewDockingComputer()
	c.Run(lib.AsStringArray(input))
	v := c.SumValuesInMemory()
	if v != 208 {
		t.Fatal(v, 208)
	}
}
