package main

import "testing"

func TestEarliestBus(t *testing.T) {
	const input = `939
7,13,x,x,59,x,31,19`

	bs := NewBusService(input)
	b := bs.FindEarliestBus()
	res := b.ID * b.Wait
	if res != 295 {
		t.Fatal(res, 295)
	}
}
