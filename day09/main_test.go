package main

import (
	"testing"
)

func TestFirstNumber(t *testing.T) {
	const input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

	x := NewXmasChecker(5, input)
	nb := x.FindFirstEncodingError()
	if nb != 127 {
		t.Fatal(nb, 127)
	}
}

func TestFindEncryptionWeakness(t *testing.T) {
	const input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

	x := NewXmasChecker(5, input)
	e := x.FindEncryptionWeakness(127)
	if e != 62 {
		t.Fatal(e, 62)
	}
}
