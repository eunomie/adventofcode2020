package main

import (
	"strings"
	"testing"
)

const (
	input = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`
)

func TestValidPasswords(t *testing.T) {
	pwdValidator := NewPasswordValidator(strings.Split(input, "\n"))
	nbValid := pwdValidator.CountValidPasswords()
	if nbValid != 2 {
		t.Fatal(nbValid, 2)
	}
}
