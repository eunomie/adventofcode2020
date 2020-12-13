package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	bs := NewBusService(lib.Input())
	b := bs.FindEarliestBus()
	res := b.ID * b.Wait
	fmt.Println("res:", res)
}

type Bus struct {
	ID   int
	Wait int
}

type BusService struct {
	timestamp int
	schedule  []int
}

func NewBusService(input string) *BusService {
	i := lib.AsStringArray(input)
	timestamp, _ := strconv.Atoi(i[0])
	schedule := []int{}
	for _, id := range strings.Split(i[1], ",") {
		if id == "x" {
			continue
		}
		busID, _ := strconv.Atoi(id)
		schedule = append(schedule, busID)
	}

	return &BusService{
		timestamp: timestamp,
		schedule:  schedule,
	}
}

func (bs *BusService) FindEarliestBus() Bus {
	wait := math.MaxInt32
	busID := 0
	for _, id := range bs.schedule {
		w := ((bs.timestamp/id)+1)*id - bs.timestamp
		if w < wait {
			wait = w
			busID = id
		}
	}
	return Bus{
		ID:   busID,
		Wait: wait,
	}
}
