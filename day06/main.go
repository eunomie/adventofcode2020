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

	fmt.Println("nb questions yes:", CountNbYes(file))
}

func CountNbYes(input io.Reader) int {
	count := 0
	questions := map[string]int{}
	nbPeople := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for _, nbAnswers := range questions {
				if nbAnswers == nbPeople {
					count++
				}
			}
			nbPeople = 0
			questions = map[string]int{}
			continue
		}
		nbPeople++
		tokens := strings.Split(line, "")
		for _, q := range tokens {
			questions[q]++
		}
	}

	for _, nbAnswers := range questions {
		if nbAnswers == nbPeople {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return count
}
