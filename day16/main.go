package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	ts := NewTicketScanner(lib.Input())
	errRate := ts.Scan()
	fmt.Println("error rate:", errRate)
	sum := ts.Departure()
	fmt.Println("departure:", sum)
}

type FieldConstraint struct {
	int1Min int
	int1Max int
	int2Min int
	int2Max int
}

func NewFieldConstraint(int1Min, int1Max, int2Min, int2Max int) *FieldConstraint {
	return &FieldConstraint{
		int1Min: int1Min,
		int1Max: int1Max,
		int2Min: int2Min,
		int2Max: int2Max,
	}
}

func (c *FieldConstraint) IsValid(val int) bool {
	return (val >= c.int1Min && val <= c.int1Max) || (val >= c.int2Min && val <= c.int2Max)
}

type TicketScanner struct {
	fields        map[string]*FieldConstraint
	yourTicket    []int
	nearbyTickets [][]int
}

func NewTicketScanner(input string) *TicketScanner {
	t := TicketScanner{
		fields:        map[string]*FieldConstraint{},
		nearbyTickets: [][]int{},
	}

	pattern := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	step := "constraint"
	for _, line := range lib.AsStringArray(input) {
		if len(line) == 0 {
			continue
		}
		if line == "your ticket:" {
			step = "your"
			continue
		}
		if line == "nearby tickets:" {
			step = "nearby"
			continue
		}
		switch step {
		case "constraint":
			match := pattern.FindStringSubmatch(line)
			int1Min, _ := strconv.Atoi(match[2])
			int1Max, _ := strconv.Atoi(match[3])
			int2Min, _ := strconv.Atoi(match[4])
			int2Max, _ := strconv.Atoi(match[5])
			t.fields[match[1]] = NewFieldConstraint(int1Min, int1Max, int2Min, int2Max)
		case "your":
			t.yourTicket = lib.AsIntArray(strings.Split(line, ","))
		case "nearby":
			t.nearbyTickets = append(t.nearbyTickets, lib.AsIntArray(strings.Split(line, ",")))
		}
	}

	return &t
}

func (t *TicketScanner) IsValid(val int) bool {
	for _, c := range t.fields {
		if c.IsValid(val) {
			return true
		}
	}
	return false
}

func (t *TicketScanner) Scan() int {
	errRate := 0

	for _, ticket := range t.nearbyTickets {
		for _, v := range ticket {
			if !t.IsValid(v) {
				errRate += v
			}
		}
	}

	return errRate
}

func (t *TicketScanner) IsValidTicket(ticket []int) bool {
	for _, v := range ticket {
		if !t.IsValid(v) {
			return false
		}
	}
	return true
}

func (t *TicketScanner) ScanFields() map[string]int {
	allFields := map[int]map[string]bool{}
	for i := range t.yourTicket {
		allFields[i] = map[string]bool{}
	}

	for _, ticket := range t.nearbyTickets {
		for i, v := range ticket {
			if !t.IsValid(v) {
				// ignore it
				// there's some duplication, but let's see performances later, if needed
				continue
			}
			for fieldName, constraint := range t.fields {
				valid := constraint.IsValid(v)
				if set, ok := allFields[i][fieldName]; ok {
					valid = valid && set
				}
				allFields[i][fieldName] = valid
			}
		}
	}

	res := map[string]int{}
	toRemove := map[string]bool{}
	for len(res) < len(allFields) {
		for _, fieldNames := range allFields {
			for k := range toRemove {
				fieldNames[k] = false
			}
		}
		toRemove = map[string]bool{}
		for i, fieldNames := range allFields {
			nbTrue, name := countAndLastTrue(fieldNames)
			if nbTrue == 1 {
				res[name] = t.yourTicket[i]
				toRemove[name] = true
			}
		}
	}

	return res
}

func (t *TicketScanner) Departure() int {
	fields := t.ScanFields()
	res := 1
	for k, v := range fields {
		if strings.HasPrefix(k, "departure") {
			res *= v
		}
	}

	return res
}

func countAndLastTrue(m map[string]bool) (int, string) {
	res := 0
	name := ""
	for n, v := range m {
		if v {
			res++
			name = n
		}
	}
	return res, name
}
