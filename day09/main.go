package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := NewXmasChecker(25, file)
	fmt.Println("first number:", x.FindFirstEncodingError())
}

type XmasChecker struct {
	reader  io.Reader
	bufsize int
}

func NewXmasChecker(bufsize int, reader io.Reader) *XmasChecker {
	return &XmasChecker{
		reader:  reader,
		bufsize: bufsize,
	}
}

func (x *XmasChecker) FindFirstEncodingError() int {
	buffer := make([]int, x.bufsize)
	bufidx := 0

	scanner := bufio.NewScanner(x.reader)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		value, _ := strconv.Atoi(line)

		if i >= x.bufsize {
			if !hasPair(value, buffer) {
				return value
			}
		}

		buffer[bufidx] = value
		bufidx = (bufidx + 1) % x.bufsize

		i++
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
