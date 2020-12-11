package main

import "testing"

func TestOccupiedSeats(t *testing.T) {
	const input = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	s := NewSeatsSimulation(input)
	finalState := s.Run()
	seats := s.CountOccupiedSeats(finalState)
	if seats != 26 {
		t.Fatal(seats, 26)
	}
}
