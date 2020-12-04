package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
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

	valid, _ := r.CountValidPassportFields()
	fmt.Println("passports with valid fields:", valid)
}

type PassportReader struct {
	passports []*Passport
}

func NewPassportReader(input io.Reader) *PassportReader {
	r := PassportReader{
		passports: []*Passport{},
	}

	passport := map[string]string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			r.passports = append(r.passports, NewPassport(passport))
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
		r.passports = append(r.passports, NewPassport(passport))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &r
}

func (r *PassportReader) CountPassportsWithMandatoryFields() int {
	validCount := 0
	r.scanPassport(func(passport *Passport) {
		if passport.HasMandatoryFields() {
			validCount++
		}
	})

	return validCount
}

func (r *PassportReader) CountValidPassportFields() (int, int) {
	validCount := 0
	invalidCount := 0
	r.scanPassport(func(passport *Passport) {
		if passport.IsValid() {
			validCount++
		} else {
			invalidCount++
		}
	})

	return validCount, invalidCount
}

func (r *PassportReader) scanPassport(callback func(passport *Passport)) {
	for _, passport := range r.passports {
		callback(passport)
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
	validators      = map[string]func(value string) bool{
		"byr": byrValidator,
		"iyr": iyrValidator,
		"eyr": eyrValidator,
		"hgt": hgtValidator,
		"hcl": hclValidator,
		"ecl": eclValidator,
		"pid": pidlValidator,
	}
)

func (p *Passport) HasMandatoryFields() bool {
	for _, field := range mandatoryFields {
		if _, present := p.fields[field]; !present {
			return false
		}
	}
	return true
}

func (p *Passport) IsValid() bool {
	for _, field := range mandatoryFields {
		value, present := p.fields[field]
		if !present {
			return false
		}
		if !validators[field](value) {
			return false
		}
	}
	return true
}

/*
byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
*/

func byrValidator(value string) bool {
	if len(value) != 4 {
		return false
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return v >= 1920 && v <= 2002
}

func iyrValidator(value string) bool {
	if len(value) != 4 {
		return false
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return v >= 2010 && v <= 2020
}

func eyrValidator(value string) bool {
	if len(value) != 4 {
		return false
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return v >= 2020 && v <= 2030
}

func hgtValidator(value string) bool {
	var unit string
	if strings.HasSuffix(value, "cm") {
		unit = "cm"
	} else if strings.HasSuffix(value, "in") {
		unit = "in"
	} else {
		return false
	}
	val := strings.TrimSuffix(value, unit)

	v, err := strconv.Atoi(val)
	if err != nil {
		return false
	}

	if unit == "cm" {
		return v >= 150 && v <= 193
	}
	if unit == "in" {
		return v >= 59 && v <= 76
	}

	return false
}

func hclValidator(value string) bool {
	pattern := regexp.MustCompile(`^#[0-9a-z]{6}$`)

	return pattern.Match([]byte(value))
}

func eclValidator(value string) bool {
	pattern := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)

	return pattern.Match([]byte(value))
}

func pidlValidator(value string) bool {
	pattern := regexp.MustCompile(`^[0-9]{9}$`)

	return pattern.Match([]byte(value))
}
