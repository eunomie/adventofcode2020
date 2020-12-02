package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	validator := NewPasswordValidator(strings.Split(string(content), "\n"))
	fmt.Println("nb valid:", validator.CountValidPasswords())
}

type PasswordValidator struct {
	passwords []*PasswordCheck
}

func NewPasswordValidator(input []string) *PasswordValidator {
	pattern := regexp.MustCompile(`^(?P<min>\d+)-(?P<max>\d+) (?P<char>\w): (?P<password>\w+)$`)
	p := []*PasswordCheck{}

	for _, line := range input {
		match := pattern.FindStringSubmatch(line)
		min, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		var char int32
		for _, c := range match[3] {
			char = c
			break
		}
		p = append(p, &PasswordCheck{
			char:          char,
			minOccurrence: min,
			maxOccurrence: max,
			password:      match[4],
		})
	}

	return &PasswordValidator{
		passwords: p,
	}
}

func (v *PasswordValidator) CountValidPasswords() int {
	count := 0
	for _, p := range v.passwords {
		if p.IsValid() {
			count++
		}
	}

	return count
}

type PasswordCheck struct {
	char          int32
	minOccurrence int
	maxOccurrence int
	password      string
}

func (c *PasswordCheck) IsValid() bool {
	count := 0
	for _, char := range c.password {
		if char != c.char {
			continue
		}
		count++
		if count > c.maxOccurrence {
			return false
		}
	}
	return count >= c.minOccurrence
}
