package main

import "testing"

func TestMath(t *testing.T) {
	testCases := []struct {
		formula string
		res     int
	}{
		{
			formula: "1 + 2 * 3 + 4 * 5 + 6",
			res:     71,
		},
		{
			formula: "1 + (2 * 3) + (4 * (5 + 6))",
			res:     51,
		},
		{
			formula: "2 * 3 + (4 * 5)",
			res:     26,
		},
		{
			formula: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			res:     437,
		},
		{
			formula: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			res:     12240,
		},
		{
			formula: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			res:     13632,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.formula, func(t *testing.T) {
			res := Compute(tc.formula)
			if res != tc.res {
				t.Error(res, tc.res)
			}
		})
	}
}
