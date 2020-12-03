package main

import "testing"

const (
	input = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
)

func TestNbTrees(t *testing.T) {
	traj := NewTrajectory(input)
	nbTrees := traj.CountTrees()
	if nbTrees != 7 {
		t.Fatal(nbTrees, 7)
	}
}

func TestAllSlopes(t *testing.T) {
	traj := NewTrajectory(input)
	nb := traj.CountTreesAllSlopes()
	if nb != 336 {
		t.Fatal(nb, 336)
	}
}
