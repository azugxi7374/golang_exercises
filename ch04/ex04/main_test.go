package main

import (
	"testing"
)

func TestShaString(t *testing.T) {
	var tests = []struct {
		input []int
        r int
		want  []int
	}{
		{[]int{}, 5, []int{}},
        {[]int{1,2,3,4,5,6}, 0, []int{1,2,3,4,5,6}},
        {[]int{1,2,3,4,5,6}, 2, []int{5,6,1,2,3,4}},
        {[]int{1,2,3,4,5,6}, 5, []int{2,3,4,5,6,1}},
	}

	for _, test := range tests {
        got := make([]int, len(test.input))
        copy(got, test.input)
        rotate(got, test.r)

        if !sliceEquals(test.want, got) {
            t.Errorf("input: %v, r: %v, want: %v, got: %v", test.input, test.r, test.want, got)
		}
	}
}

func sliceEquals(a []int, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, _ := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
