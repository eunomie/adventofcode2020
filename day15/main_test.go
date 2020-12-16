package main

import (
	"fmt"
	"testing"
)

func Test2020thTurn(t *testing.T) {
	testCases := []struct {
		input []int
		res   int
	}{
		{
			input: []int{0, 3, 6},
			res:   436,
		},
		{
			input: []int{1, 3, 2},
			res:   1,
		},
		{
			input: []int{2, 1, 3},
			res:   10,
		},
		{
			input: []int{1, 2, 3},
			res:   27,
		},
		{
			input: []int{2, 3, 1},
			res:   78,
		},
		{
			input: []int{3, 2, 1},
			res:   438,
		},
		{
			input: []int{3, 1, 2},
			res:   1836,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := Get2020thTurn(tc.input)
			if res != tc.res {
				t.Error(res, tc.res)
			}
		})
	}
}

func Test30000000thTurn(t *testing.T) {
	testCases := []struct {
		input []int
		res   int
	}{
		{
			input: []int{0, 3, 6},
			res:   175594,
		},
		{
			input: []int{1, 3, 2},
			res:   2578,
		},
		{
			input: []int{2, 1, 3},
			res:   3544142,
		},
		{
			input: []int{1, 2, 3},
			res:   261214,
		},
		{
			input: []int{2, 3, 1},
			res:   6895259,
		},
		{
			input: []int{3, 2, 1},
			res:   18,
		},
		{
			input: []int{3, 1, 2},
			res:   362,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := Get30000000thTurn(tc.input)
			if res != tc.res {
				t.Error(res, tc.res)
			}
		})
	}
}
