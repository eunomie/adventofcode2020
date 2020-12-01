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
	v1, v2 := input.GetPair()

	fmt.Printf("pair is %d - %d\n", v1, v2)
	fmt.Println(v1 * v2)
}

type Input struct {
	values map[int]bool
}

func NewInput(inputContent string) *Input {
	input := Input{
		values: map[int]bool{},
	}
	for _, line := range strings.Split(inputContent, "\n") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		input.values[val] = true
	}

	return &input
}

func (i *Input) GetPair() (int, int) {
	for val := range i.values {
		missing := 2020 - val
		if _, ok := i.values[missing]; ok {
			return val, missing
		}
	}
	return 0, 0
}
