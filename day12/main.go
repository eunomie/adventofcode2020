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
	waypoint := Position{10, 1}
	pos := Position{0, 0}

	for _, direction := range f.directions {
		value := direction[1:]
		v, _ := strconv.Atoi(value)
		switch direction[0] {
		case 'R', 'L':
			angle := v
			if direction[0] == 'R' {
				angle = 360 - angle
			}
			waypoint = Position{
				E: waypoint.E*cos[angle] - waypoint.N*sin[angle],
				N: waypoint.E*sin[angle] + waypoint.N*cos[angle],
			}
		case 'N':
			waypoint.N += v
		case 'S':
			waypoint.N -= v
		case 'E':
			waypoint.E += v
		case 'W':
			waypoint.E -= v
		case 'F':
			pos.N += v * waypoint.N
			pos.E += v * waypoint.E
		}
	}
	return pos
}

var (
	cos = map[int]int{
		0:   1,
		90:  0,
		180: -1,
		270: 0,
	}
	sin = map[int]int{
		0:   0,
		90:  1,
		180: 0,
		270: -1,
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
