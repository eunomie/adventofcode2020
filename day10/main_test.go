package main

import (
	"fmt"
	"testing"
)

func TestAdapters(t *testing.T) {
	testCases := []struct {
		input string
		res   int
	}{
		{
			input: `16
10
15
5
1
11
7
19
6
12
4`,
			res: 7 * 5,
		},
		{
			input: `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`,
			res: 22 * 10,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			a := NewAdapterSequence(tc.input)
			seqValue := a.GetAdapterDifferences()
			if seqValue != tc.res {
				t.Fatal(seqValue, tc.res)
			}
		})
	}
}
