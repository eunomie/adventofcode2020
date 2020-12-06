package main

import (
	"strings"
	"testing"
)

func TestNbYes(t *testing.T) {
	const input = `abc

a
b
c

ab
ac

a
a
a
a

b`

	nbYes := CountNbYes(strings.NewReader(input))

	if nbYes != 6 {
		t.Fatal(nbYes, 6)
	}
}
