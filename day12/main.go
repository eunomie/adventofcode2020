package main

import (
	"fmt"
	"strconv"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	f := NewFerry(lib.AsStringArray(lib.Input()))
	pos := f.Move()
	dist := pos.ManhattanDistance()
	fmt.Println("Manhattan Distance:", dist)
}

type Ferry struct {
	directions []string
}

func NewFerry(directions []string) *Ferry {
	return &Ferry{
		directions: directions,
	}
}

func (f *Ferry) Move() Position {
	orientation := OrientationE
	orientationIdx := 0
	pos := Position{0, 0}

	for _, direction := range f.directions {
		value := direction[1:]
		v, _ := strconv.Atoi(value)
		switch direction[0] {
		case 'R':
			orientationIdx = (orientationIdx + 1*v/90) % 4
			orientation = OrientationsR[orientationIdx]
		case 'L':
			orientationIdx = (orientationIdx + 3*v/90) % 4
			orientation = OrientationsR[orientationIdx]
		case 'N':
			pos.N += v
		case 'S':
			pos.N -= v
		case 'E':
			pos.E += v
		case 'W':
			pos.E -= v
		case 'F':
			pos.N += v * orientation.N
			pos.E += v * orientation.E
		}
	}
	return pos
}

type Orientation struct {
	E int
	N int
}

var (
	OrientationE  = Orientation{1, 0}
	OrientationN  = Orientation{0, 1}
	OrientationW  = Orientation{-1, 0}
	OrientationS  = Orientation{0, -1}
	OrientationsR = []Orientation{
		OrientationE, OrientationS, OrientationW, OrientationN,
	}
)

type Position struct {
	E int
	N int
}

func (p *Position) ManhattanDistance() int {
	return absInt(p.E) + absInt(p.N)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
