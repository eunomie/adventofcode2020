package main

import "testing"

func TestActiveAfter6Cycles(t *testing.T) {
	const input = `.#.
..#
###`

	cc := NewConwayCubes(input)
	cc.Run(6)
	res := cc.ActiveStates()
	if res != 112 {
		t.Fatal(res, 112)
	}
}
