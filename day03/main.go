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
	fmt.Println("nb trees:", traj.CountTrees())
	fmt.Println("all slopes:", traj.CountTreesAllSlopes())
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
	return t.countTrees(3, 1)
}

func (t *Trajectory) countTrees(incX, incY int) int {
	x := 0
	y := 0
	nbTrees := 0

	for {
		x = (x + incX) % t.nbColumns
		y += incY
		if y >= t.nbLines {
			break
		}
		if t.trees[y][x] == "#" {
			nbTrees++
		}
	}

	return nbTrees
}

func (t *Trajectory) CountTreesAllSlopes() int {
	return t.countTrees(1, 1) *
		t.countTrees(3, 1) *
		t.countTrees(5, 1) *
		t.countTrees(7, 1) *
		t.countTrees(1, 2)
}
