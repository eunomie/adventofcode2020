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
			if seats[i][j] == "#" && occupied >= 5 {
				newSeats[i][j] = "L"
			}
		}
	}
	return newSeats
}

func (s *SeatsSimulation) occupiedSeats(seats [][]string, x, y int) int {
	nbOccupied := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				if s.occupiedSeat(seats, x, y, i, j) {
					nbOccupied++
				}
			}
		}
	}
	return nbOccupied
}

func (s *SeatsSimulation) occupiedSeat(seats [][]string, x, y, xOffset, yOffset int) bool {
	i := x + xOffset
	j := y + yOffset
	for i >= 0 && i < s.lines && j >= 0 && j < s.columns {
		if !(i == x && j == y) {
			if seats[i][j] == "#" {
				return true
			}
			if seats[i][j] == "L" {
				return false
			}
		}
		i += xOffset
		j += yOffset
	}
	return false
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
