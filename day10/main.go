package main

import (
	"fmt"
	"sort"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	a := NewAdapterSequence(lib.Input())
	fmt.Println("seq:", a.GetAdapterDifferences())
}

type AdapterSequence struct {
	input []int
}

func NewAdapterSequence(input string) *AdapterSequence {
	return &AdapterSequence{
		input: lib.AsIntArray(lib.AsStringArray(input)),
	}
}

func (a *AdapterSequence) GetAdapterDifferences() int {
	s := make([]int, len(a.input))
	copy(s, a.input)

	sort.Ints(s)

	oneJoltDiffs := 0
	threeJoltDiffs := 1 // +1 because it always ends up with a +3 jolts
	current := 0
	for _, v := range s {
		switch v - current {
		case 1:
			oneJoltDiffs++
		case 3:
			threeJoltDiffs++
		}
		current = v
	}
	return oneJoltDiffs * threeJoltDiffs
}
