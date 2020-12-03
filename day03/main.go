package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	traj := NewTrajectory(string(content))
	fmt.Println("nbTrees:", traj.CountTrees())
}

type Trajectory struct {
	trees     [][]string
	nbLines   int
	nbColumns int
}

func NewTrajectory(input string) *Trajectory {
	traj := Trajectory{
		trees: [][]string{},
	}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		places := strings.Split(line, "")
		traj.trees = append(traj.trees, places)
	}

	traj.nbLines = len(traj.trees)
	traj.nbColumns = len(traj.trees[0])

	return &traj
}

func (t *Trajectory) CountTrees() int {
	x := 0
	y := 0
	nbTrees := 0

	for y < t.nbLines-1 {
		x = (x + 3) % t.nbColumns
		y += 1
		if t.trees[y][x] == "#" {
			nbTrees++
		}
	}

	return nbTrees
}
