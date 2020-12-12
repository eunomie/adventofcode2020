package main

import (
	"testing"

	"github.com/eunomie/adventofcode2020/lib"
)

func TestFerryNavigation(t *testing.T) {
	const input = `F10
N3
F7
R90
F11`

	f := NewFerry(lib.AsStringArray(input))
	pos := f.Move()
	dist := pos.ManhattanDistance()
	if dist != 286 {
		t.Fatal(dist, 286)
	}
}
