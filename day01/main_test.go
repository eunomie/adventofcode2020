package main

import "testing"

const (
	content = `1721
979
366
299
675
1456`
)

func TestPair(t *testing.T) {
	input := NewInput(content)
	pair := input.GetPairProduct()
	if pair != 1721*299 {
		t.Fatal(pair, 1721*289)
	}
}

func TestTriple(t *testing.T) {
	input := NewInput(content)
	triple := input.GetTripleProduct()
	if triple != 241861950 {
		t.Fatal(triple, 241861950)
	}
}
