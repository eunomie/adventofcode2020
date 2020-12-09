package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	x := NewXmasChecker(25, string(content))
	nb := x.FindFirstEncodingError()
	fmt.Println("first number:", nb)
	fmt.Println("encryption weakness:", x.FindEncryptionWeakness(nb))
}

type XmasChecker struct {
	sequence []int
	bufsize  int
}

func NewXmasChecker(bufsize int, content string) *XmasChecker {
	lines := strings.Split(content, "\n")
	sequence := make([]int, len(lines))
	for i, line := range lines {
		v, _ := strconv.Atoi(line)
		sequence[i] = v
	}
	return &XmasChecker{
		sequence: sequence,
		bufsize:  bufsize,
	}
}

func (x *XmasChecker) FindFirstEncodingError() int {
	buffer := make([]int, x.bufsize)
	bufidx := 0

	for i, value := range x.sequence {
		if i >= x.bufsize {
			if !hasPair(value, buffer) {
				return value
			}
		}

		buffer[bufidx] = value
		bufidx = (bufidx + 1) % x.bufsize
	}
	return 0
}

func hasPair(value int, buffer []int) bool {
	for i, first := range buffer {
		missing := value - first
		for _, second := range buffer[i:] {
			if first == second {
				continue
			}
			if second == missing {
				return true
			}
		}
	}
	return false
}

func (x *XmasChecker) FindEncryptionWeakness(nb int) int {
	sum := 0
	min := math.MaxInt32
	max := 0
	idx := 0
	startRangeIdx := 0
	l := len(x.sequence)
	for {
		if idx == l {
			break
		}
		value := x.sequence[idx]
		sum += value
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}

		if sum == nb {
			return min + max
		}
		if sum > nb {
			sum = 0
			min = math.MaxInt32
			max = 0
			startRangeIdx++
			idx = startRangeIdx
			continue
		}
		idx++
	}

	fmt.Println(l, idx, startRangeIdx)

	return 0
}
