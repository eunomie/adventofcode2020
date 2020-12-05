package main

import (
	"testing"
)

func TestBoardingPassPosition(t *testing.T) {
	testCases := []struct {
		input string
		row   int
		col   int
	}{
		{
			input: "FBFBBFFRLR",
			row:   44,
			col:   5,
		},
		{
			input: "BFFFBBFRRR",
			row:   70,
			col:   7,
		},
		{
			input: "FFFBBBFRRR",
			row:   14,
			col:   7,
		},
		{
			input: "BBFFBBFRLL",
			row:   102,
			col:   4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			b := NewBoardingPass(tc.input)
			if b.Row != tc.row {
				t.Error(b.Row, tc.row)
			}
			if b.Col != tc.col {
				t.Error(b.Col, tc.col)
			}
		})
	}
}

func TestHightSeatID(t *testing.T) {
	const input = `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

	f := NewBoardingPassFinder(input)
	highestSeatID := f.GetHigestSeatID()
	if highestSeatID != 820 {
		t.Fatal(highestSeatID, 820)
	}
}
