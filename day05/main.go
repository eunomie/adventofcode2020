package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	f := NewBoardingPassFinder(string(content))
	fmt.Println("Hight boarding pass:", f.GetHigestSeatID())
	fmt.Println("Empty seat:", f.GetEmptySeat())
}

type BoardingPassFinder struct {
	BoardingPass []*BoardingPass
}

func NewBoardingPassFinder(input string) *BoardingPassFinder {
	f := BoardingPassFinder{BoardingPass: []*BoardingPass{}}
	for _, line := range strings.Split(input, "\n") {
		b := NewBoardingPass(line)
		f.BoardingPass = append(f.BoardingPass, b)
	}
	return &f
}

func (f *BoardingPassFinder) GetHigestSeatID() int {
	highest := 0
	for _, b := range f.BoardingPass {
		seatID := b.SeatID()
		if seatID > highest {
			highest = seatID
		}
	}
	return highest
}

func (f *BoardingPassFinder) GetEmptySeat() int {
	seats := map[int]bool{}
	for _, b := range f.BoardingPass {
		seats[b.SeatID()] = true
	}
	for row := 1; row < 126; row++ {
		for col := 0; col < 8; col++ {
			seatID := row*8 + col
			if _, present := seats[seatID]; !present {
				_, prev := seats[seatID-1]
				_, next := seats[seatID+1]
				if prev && next {
					return seatID
				}
			}
		}
	}
	return 0
}

type BoardingPass struct {
	RowDef string
	ColDef string
	Row    int
	Col    int
}

func NewBoardingPass(value string) *BoardingPass {
	b := &BoardingPass{
		RowDef: value[:7],
		ColDef: value[7:],
	}
	b.ComputeSeatPosition()
	return b
}

func (b *BoardingPass) ComputeSeatPosition() {
	b.Row = computePosition(b.RowDef, 0, 127)
	b.Col = computePosition(b.ColDef, 0, 8)
}

func (b *BoardingPass) SeatID() int {
	return b.Row*8 + b.Col
}

func computePosition(val string, l, u int) int {
	lRange := l
	uRange := u
	for _, v := range val {
		if v == 'F' || v == 'L' {
			uRange = (lRange + uRange - 1) / 2
		}
		if v == 'B' || v == 'R' {
			lRange = (lRange + uRange + 1) / 2
		}
	}
	return lRange
}
