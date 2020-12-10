package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	a := NewAdapterSequence(lib.Input())
	fmt.Println("seq:", a.GetAdapterDifferences())
	fmt.Println("combinations:", a.GetCombinations())
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
	seq := a.getDefaultSequence()

	oneJoltDiffs := 0
	threeJoltDiffs := 0
	for _, v := range seq {
		if v == 1 {
			oneJoltDiffs++
		}
		if v == 3 {
			threeJoltDiffs++
		}
	}

	return oneJoltDiffs * threeJoltDiffs
}

func (a *AdapterSequence) getDefaultSequence() []int {
	s := make([]int, len(a.input))
	copy(s, a.input)

	seq := make([]int, len(a.input)+1)

	sort.Ints(s)

	current := 0
	for i, v := range s {
		seq[i] = v - current
		current = v
	}
	seq[len(seq)-1] = 3
	return seq
}

func (a *AdapterSequence) GetCombinations() int {
	seq := a.getDefaultSequence()

	combinations := 1
	nbOnes := 0
	addCombinations := func() {
		switch nbOnes {
		case 1:
		case 2:
			combinations = combinations * 2
		default:
			combinations = combinations * (1 + 3*powInt(2, nbOnes-3))
		}
	}
	for _, v := range seq {
		if v == 1 {
			nbOnes++
		}
		if v == 3 {
			addCombinations()
			nbOnes = 0
		}
	}

	addCombinations()

	return combinations
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
