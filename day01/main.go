package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	input := NewInput(string(content))

	fmt.Println("pair:", input.GetPairProduct())
	fmt.Println("triple:", input.GetTripleProduct())
}

type Input struct {
	vals  []int
	pairs map[int]int
}

func NewInput(inputContent string) *Input {
	input := Input{
		pairs: map[int]int{},
	}
	lines := strings.Split(inputContent, "\n")
	vals := make([]int, len(lines))
	for i, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		vals[i] = val
	}

	for i := 0; i < len(vals); i++ {
		for j := i + 1; j < len(vals); j++ {
			input.pairs[vals[i]+vals[j]] = vals[i] * vals[j]
		}
	}
	input.vals = vals

	return &input
}

func (i *Input) GetPairProduct() int {
	return i.pairs[2020]
}

func (i *Input) GetTripleProduct() int {
	for _, val := range i.vals {
		missing := 2020 - val
		if pair, ok := i.pairs[missing]; ok {
			return pair * val
		}
	}
	return 0
}
