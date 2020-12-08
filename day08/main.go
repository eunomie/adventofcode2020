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
	_ = b.FixAndRun()
	fmt.Println("fixed:", b.Acc)
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
	_, acc := run(b.instructions)
	b.Acc = acc
}

func run(instructions []string) (bool, int) {
	acc := 0
	nextInstruction := len(instructions)
	idx := 0
	alreadyRunIdx := map[int]bool{}
	eof := false

	for {
		if eof {
			break
		}
		if _, ran := alreadyRunIdx[idx]; ran {
			break
		}
		alreadyRunIdx[idx] = true
		instruction := instructions[idx]

		if strings.HasPrefix(instruction, "nop") {
			idx++
			eof = idx == nextInstruction
			continue
		}

		if strings.HasPrefix(instruction, "acc") {
			incStr := strings.TrimPrefix(instruction, "acc ")
			inc, _ := strconv.Atoi(incStr)
			acc += inc
			idx++
			eof = idx == nextInstruction
			continue
		}

		if strings.HasPrefix(instruction, "jmp") {
			jmpStr := strings.TrimPrefix(instruction, "jmp ")
			jmp, _ := strconv.Atoi(jmpStr)
			idx += jmp
			eof = idx == nextInstruction
			continue
		}
	}
	return eof, acc
}

func (b *BootSequence) FixAndRun() bool {
	fixInstructionAndRun := func(from []string, idx int, instruction, before, after string) (bool, int) {
		instructions := make([]string, len(b.instructions))
		copy(instructions, b.instructions)
		newInstruction := after + strings.TrimPrefix(instruction, before)
		instructions[idx] = newInstruction
		return run(instructions)

	}

	for idx, instruction := range b.instructions {
		if strings.HasPrefix(instruction, "nop") {
			eof, acc := fixInstructionAndRun(b.instructions, idx, instruction, "nop ", "jmp ")
			if eof {
				b.Acc = acc
				return true
			}
		}
		if strings.HasPrefix(instruction, "jmp") {
			eof, acc := fixInstructionAndRun(b.instructions, idx, instruction, "jmp ", "nop ")
			if eof {
				b.Acc = acc
				return true
			}
		}
	}
	return false
}
