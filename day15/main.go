package main

import (
	"fmt"
	"strings"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	fmt.Println("2020:", Get2020thTurn(lib.AsIntArray(strings.Split(lib.Input(), ","))))
	fmt.Println("30000000:", Get30000000thTurn(lib.AsIntArray(strings.Split(lib.Input(), ","))))
}

func Get2020thTurn(input []int) int {
	return getTurn(input, 2020)
}

func Get30000000thTurn(input []int) int {
	return getTurn(input, 30000000)
}

func getTurn(input []int, limit int) int {
	seq := []int{}
	numbers := map[int][]int{}

	for i, v := range input {
		numbers[v] = []int{i + 1}
		seq = append(seq, v)
	}

	lastNum := input[len(input)-1]

	for i := len(input) + 1; i <= limit; i++ {
		n, exist := numbers[lastNum]
		nextNum := 0
		if exist && len(n) > 1 {
			nextNum = n[len(n)-1] - n[len(n)-2]
		}
		if _, ok := numbers[nextNum]; !ok {
			numbers[nextNum] = []int{}
		}
		numbers[nextNum] = append(numbers[nextNum], i)
		seq = append(seq, nextNum)
		lastNum = nextNum
	}

	return lastNum
}
