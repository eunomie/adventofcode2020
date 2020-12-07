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

	r := NewBagsRulesReader(file)
	fmt.Println("number of bags:", r.CountBagsForColor("shiny gold"))
	fmt.Println("number of bags inside:", r.CountBagsInside("shiny gold"))
}

func read(reader io.Reader, callback func(line string)) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		callback(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type BagsRulesReader struct {
	Bags map[string]*Bag
}

func NewBagsRulesReader(reader io.Reader) *BagsRulesReader {
	r := BagsRulesReader{
		Bags: map[string]*Bag{},
	}

	getOrInitBag := func(color string) *Bag {
		bag, ok := r.Bags[color]
		if !ok {
			bag = NewBag(color)
			r.Bags[color] = bag
		}
		return bag
	}

	updateLinks := func(baseBag, childBag *Bag, nbChildren int) {
		childBag.Ancestors = append(childBag.Ancestors, baseBag)
		baseBag.Children[childBag] = nbChildren
		r.Bags[childBag.Color] = childBag
	}

	read(reader, func(line string) {
		pattern := regexp.MustCompile(`^([\w ]+) bags contain,?(.*).$`)
		childPattern := regexp.MustCompile(`^ (\d) ([\w ]+) bags?`)
		match := pattern.FindStringSubmatch(line)

		bagColor := match[1]
		bag := getOrInitBag(bagColor)

		if match[2] != " no other bags" {
			tokens := strings.Split(match[2], ",")
			for _, t := range tokens {
				match := childPattern.FindStringSubmatch(t)
				childBag := getOrInitBag(match[2])
				nbChildren, _ := strconv.Atoi(match[1])
				updateLinks(bag, childBag, nbChildren)
			}
		}
	})

	return &r
}

func (r *BagsRulesReader) countBagsForColor(color string) map[string]bool {
	bag := r.Bags[color]
	colors := map[string]bool{}
	for _, ancestor := range bag.Ancestors {
		colors[ancestor.Color] = true
		otherColors := r.countBagsForColor(ancestor.Color)
		for c := range otherColors {
			colors[c] = true
		}
	}
	return colors
}

func (r *BagsRulesReader) CountBagsForColor(color string) int {
	return len(r.countBagsForColor(color))
}

func (r *BagsRulesReader) CountBagsInside(color string) int {
	bag := r.Bags[color]
	count := 0
	for childBag, nbChildren := range bag.Children {
		count += nbChildren
		count += nbChildren * r.CountBagsInside(childBag.Color)
	}
	return count
}

type Bag struct {
	Color     string
	Ancestors []*Bag
	Children  map[*Bag]int
}

func NewBag(color string) *Bag {
	return &Bag{
		Color:     color,
		Ancestors: []*Bag{},
		Children:  map[*Bag]int{},
	}
}
