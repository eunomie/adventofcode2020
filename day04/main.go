package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := NewPassportReader(file)
	fmt.Println("valid passports:", r.CountPassportsWithMandatoryFields())
}

type PassportReader struct {
	input io.Reader
}

func NewPassportReader(input io.Reader) *PassportReader {
	return &PassportReader{
		input: input,
	}
}

func (r *PassportReader) CountPassportsWithMandatoryFields() int {
	validCount := 0
	r.scanPassport(r.input, func(passport *Passport) {
		if passport.HasMandatoryFields() {
			validCount++
		}
	})

	return validCount
}

func (r *PassportReader) scanPassport(reader io.Reader, callback func(passport *Passport)) {
	passport := map[string]string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			callback(NewPassport(passport))
			passport = map[string]string{}
			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			tokens := strings.Split(field, ":")
			passport[tokens[0]] = tokens[1]
		}
	}
	if len(passport) != 0 {
		callback(NewPassport(passport))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Passport struct {
	fields map[string]string
}

func NewPassport(fields map[string]string) *Passport {
	return &Passport{
		fields: fields,
	}
}

var (
	mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
)

func (p *Passport) HasMandatoryFields() bool {
	for _, field := range mandatoryFields {
		if _, present := p.fields[field]; !present {
			return false
		}
	}
	return true
}
