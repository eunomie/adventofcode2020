package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	b := NewBootSequence(string(content))
	b.Run()
	fmt.Println("acc:", b.Acc)
}

type BootSequence struct {
	instructions []string
	Acc          int
}

func NewBootSequence(input string) *BootSequence {
	return &BootSequence{
		instructions: strings.Split(input, "\n"),
	}
}

func (b *BootSequence) Run() {
	b.Acc = 0
	idx := 0
	alreadyRunIdx := map[int]bool{}

	for {
		if _, ran := alreadyRunIdx[idx]; ran {
			break
		}
		alreadyRunIdx[idx] = true
		instruction := b.instructions[idx]

		if strings.HasPrefix(instruction, "nop") {
			idx++
			continue
		}

		if strings.HasPrefix(instruction, "acc") {
			incStr := strings.TrimPrefix(instruction, "acc ")
			inc, _ := strconv.Atoi(incStr)
			b.Acc += inc
			idx++
			continue
		}

		if strings.HasPrefix(instruction, "jmp") {
			jmpStr := strings.TrimPrefix(instruction, "jmp ")
			jmp, _ := strconv.Atoi(jmpStr)
			idx += jmp
			continue
		}
	}
}
