package main

import (
	"fmt"
	"strconv"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	fomulas := lib.AsStringArray(lib.Input())
	sum := 0
	for _, formula := range fomulas {
		sum += Compute(formula)
	}

	fmt.Println("sum:", sum)
}

func Compute(formula string) int {
	output := []int{}
	operators := []string{}

	var l, r, v int
	var o string

	apply := func() {
		l, r, output = output[len(output)-2], output[len(output)-1], output[:len(output)-2]
		o, operators = operators[len(operators)-1], operators[:len(operators)-1]
		if o == "+" {
			v = l + r
		} else if o == "*" {
			v = l * r
		}
		output = append(output, v)
	}

	for _, c := range formula {
		if c == ' ' {
			continue
		}
		switch c {
		case '+':
			operators = append(operators, string(c))
		case '*':
			for len(operators) > 0 && operators[len(operators)-1] == "+" {
				apply()
			}
			operators = append(operators, string(c))
		case '(':
			operators = append(operators, string(c))
		case ')':
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				apply()
			}
			operators = operators[:len(operators)-1]
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v, _ = strconv.Atoi(string(c))
			output = append(output, v)
		}
	}

	for len(output) > 1 {
		apply()
	}

	return output[0]
}
