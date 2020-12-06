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
	questions := map[string]bool{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			count += len(questions)
			questions = map[string]bool{}
			continue
		}
		tokens := strings.Split(line, "")
		for _, q := range tokens {
			questions[q] = true
		}
	}

	count += len(questions)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return count
}
