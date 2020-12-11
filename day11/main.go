package main

import (
	"fmt"
	"strings"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	s := NewSeatsSimulation(lib.Input())
	finalState := s.Run()
	fmt.Println("occupied seats:", s.CountOccupiedSeats(finalState))
}

type SeatsSimulation struct {
	initial [][]string
	lines   int
	columns int
}

func NewSeatsSimulation(input string) *SeatsSimulation {
	lines := lib.AsStringArray(input)
	initial := make([][]string, len(lines))
	for i, line := range lines {
		row := strings.Split(line, "")
		initial[i] = row
	}
	return &SeatsSimulation{
		initial: initial,
		lines:   len(initial),
		columns: len(initial[0]),
	}
}

func (s *SeatsSimulation) Run() [][]string {
	seats := s.initial
	for {
		newSeats := s.run(seats)
		if s.noChange(seats, newSeats) {
			return newSeats
		}
		seats = newSeats
	}
}

func (s *SeatsSimulation) run(seats [][]string) [][]string {
	newSeats := copyMatrix(seats)
	for i := 0; i < s.lines; i++ {
		for j := 0; j < s.columns; j++ {
			if seats[i][j] == "." {
				continue
			}
			occupied := s.occupiedSeats(seats, i, j)
			if seats[i][j] == "L" && occupied == 0 {
				newSeats[i][j] = "#"
				continue
			}
			if seats[i][j] == "#" && occupied >= 4 {
				newSeats[i][j] = "L"
			}
		}
	}
	return newSeats
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *SeatsSimulation) occupiedSeats(seats [][]string, x, y int) int {
	nbOccupied := 0
	for i := maxInt(0, x-1); i <= minInt(s.lines-1, x+1); i++ {
		for j := maxInt(0, y-1); j <= minInt(s.columns-1, y+1); j++ {
			if !(i == x && j == y) && seats[i][j] == "#" {
				nbOccupied++
			}
		}
	}
	return nbOccupied
}

func (s *SeatsSimulation) noChange(seats, newSeats [][]string) bool {
	for i := 0; i < s.lines; i++ {
		for j := 0; j < s.columns; j++ {
			if seats[i][j] != newSeats[i][j] {
				return false
			}
		}
	}
	return true
}

func (s *SeatsSimulation) CountOccupiedSeats(seats [][]string) int {
	nbSeats := 0
	for i := 0; i < s.lines; i++ {
		for j := 0; j < s.columns; j++ {
			if seats[i][j] == "#" {
				nbSeats++
			}
		}
	}
	return nbSeats
}

func copyMatrix(matrix [][]string) [][]string {
	res := make([][]string, len(matrix))
	for i, el := range matrix {
		row := make([]string, len(el))
		copy(row, el)
		res[i] = row
	}

	return res
}
