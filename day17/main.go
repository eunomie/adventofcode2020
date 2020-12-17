package main

import (
	"fmt"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	cc := NewConwayCubes(lib.Input())
	cc.Run(6)
	res := cc.ActiveStates()
	fmt.Println("Active states:", res)
}

type ConwayCubes struct {
	l     int
	state map[int]map[int]map[int]bool
}

func NewConwayCubes(input string) *ConwayCubes {
	inputLines := lib.AsStringArray(input)
	cc := ConwayCubes{
		l:     len(inputLines),
		state: map[int]map[int]map[int]bool{},
	}
	z := 0
	cc.state[z] = map[int]map[int]bool{}
	for y, line := range inputLines {
		cc.state[z][y] = map[int]bool{}
		for x, v := range line {
			cc.state[z][y][x] = v == '#'
		}
	}

	return &cc
}

func (cc *ConwayCubes) Run(nbCycles int) {
	cc.printState()
	for i := 0; i < nbCycles; i++ {
		c := map[int]map[int]map[int]bool{}
		for z := -12; z < 12; z++ {
			c[z] = map[int]map[int]bool{}
			for y := -10; y < 10+cc.l; y++ {
				c[z][y] = map[int]bool{}
				for x := -10; x < 10+cc.l; x++ {
					nbActive := cc.countActiveNeighbours(x, y, z)
					c[z][y][x] = nbActive == 3 || (cc.state[z][y][x] && nbActive == 2)
				}
			}
		}
		cc.state = c
		fmt.Println("CYCLE", i)
		cc.printState()
	}
}

func (cc *ConwayCubes) printState() {
	p := false
	for z := -12; z < 12; z++ {
		s := []string{}
		for y := -10; y < 10+cc.l; y++ {
			l := ""
			for x := -10; x < 10+cc.l; x++ {
				if cc.state[z][y][x] {
					p = true
					l += "#"
				} else {
					l += "."
				}
			}
			s = append(s, l)
		}
		if p {
			for _, l := range s {
				fmt.Println(l)
			}
			fmt.Println()
		}
		p = false
	}
}

func (cc *ConwayCubes) countActiveNeighbours(x, y, z int) int {
	active := 0
	for zz := z - 1; zz <= z+1; zz++ {
		for yy := y - 1; yy <= y+1; yy++ {
			for xx := x - 1; xx <= x+1; xx++ {
				if zz == z && yy == y && xx == x {
					continue
				}
				if cc.isActive(xx, yy, zz) {
					active++
				}
			}
		}
	}
	return active
}

func (cc *ConwayCubes) isActive(x, y, z int) bool {
	if s, ok := cc.state[z]; ok {
		if row, ok := s[y]; ok {
			if cell, ok := row[x]; ok {
				return cell
			}
		}
	}
	return false
}

func (cc *ConwayCubes) ActiveStates() int {
	nbActive := 0
	for _, s := range cc.state {
		for _, line := range s {
			for _, cell := range line {
				if cell {
					nbActive++
				}
			}
		}
	}
	return nbActive
}
