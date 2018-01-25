package intset

import (
    "testing"
)

func TestIntSet(t *testing.T) {
    tests := []struct {
        adding []int
        removing []int
        want string
        l int
    }{
        {[]int{1, 5, 13, 145}, []int{13, 5, 6, 5}, "{1 145}", 2},
        {[]int{}, []int{}, "{}", 0},
    }

    for _, test := range tests {
        var set1 IntSet

        // add
        for _, a := range test.adding {
            set1.Add(a)
        }
        // remove
        for _, r := range test.removing {
            set1.Remove(r)
        }

        // copy
        set2 := set1.Copy()

        set1.Clear()

        if set1.String() != "{}" {
            t.Errorf("set clear failed. got: %v", set1.String())
        }
        if set2.String() != test.want {
            t.Errorf("set results failed. adding: %v, removing: %v, want: %v, got: %v",
            test.adding, test.removing, test.want, set2.String())
        }
        if set2.Len() != test.l {
            t.Errorf("set length failed. adding: %v, removing: %v, want: %v, got: %v",
            test.adding, test.removing, test.l, set2.Len())
        }
    }
}
