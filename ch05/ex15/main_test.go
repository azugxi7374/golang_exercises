package main

import (
    "testing"
)

func TestMaxMin(t *testing.T){
    var tests = []struct {
        input []int
        wants []int
    }{
        {[]int{3, -2, 4, 1}, []int{4, -2, 4, -2}},
        {[]int{}, []int{0, 0}},
    }

    for _, test := range tests {
        for idx, want := range test.wants {
            var got int
            switch idx {
            case 0:
                got = max0(test.input...)
            case 1:
                got = min0(test.input...)
            case 2:
                got = max1(test.input[0], test.input[1:]...)
            case 3:
                got = min1(test.input[0], test.input[1:]...)
            }

            if want != got {
                t.Errorf("Failed. func: %v, input: %v, want: %v, got: %v",
                []string{"max0", "min0", "max1", "min1"}[idx],
                test.input, want, got)
            }
        }
    }
}
