package intset

import (
    "testing"
)

func TestIntSet(t *testing.T) {
    tests := []struct {
        input []int
    }{
        {[]int{1, 44, 123, 678}},
        {[]int{}},
    }

    for _, test := range tests {
        set := &IntSet{[]uint64{}}
        set.AddAll(test.input...)

        got := set.Elems()

        if !sameSlice(got, test.input) {
            t.Errorf("elems() failed. input: %v, got: %v",
            test.input, got)
        }
    }
}
func sameSlice(a []int, b []int) bool{
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
