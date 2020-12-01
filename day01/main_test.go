package main

import "testing"

func TestPair(t *testing.T) {
	content := `1721
979
366
299
675
1456`
	input := NewInput(content)
	v1, v2 := input.GetPair()
	if v1*v2 != 1721*299 {
		t.Fatal(v1, v2)
	}
}
