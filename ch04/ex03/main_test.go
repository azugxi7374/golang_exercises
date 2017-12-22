package main

import (
	"testing"
)

func TestShaString(t *testing.T) {
	var tests = []struct {
		input [SIZE]int
		want  [SIZE]int
	}{
		{[SIZE]int{1, 1, 1, 1, 1}, [SIZE]int{1, 1, 1, 1, 1}},
		{[SIZE]int{1, 2, 3, 4, 5}, [SIZE]int{5, 4, 3, 2, 1}},
	}

	for _, test := range tests {
        got := test.input
		reverse(&got)
		if got != test.want {
			t.Errorf("input: %v, want: %v, got: %v", test.input, test.want, got)
		}
	}
}
