package main

import "testing"

func TestAcc(t *testing.T) {
	const input = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

	b := NewBootSequence(input)
	b.Run()
	if b.Acc != 5 {
		t.Fatal(b.Acc, 5)
	}
}
