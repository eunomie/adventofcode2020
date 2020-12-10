package lib

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Input() string {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func AsStringArray(input string) []string {
	return strings.Split(input, "\n")
}

func AsIntArray(input []string) []int {
	ints := make([]int, len(input))
	for i, value := range input {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = v
	}

	return ints
}
