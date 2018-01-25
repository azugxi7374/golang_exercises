package intset

import (
    "testing"
)

func TestIntSet(t *testing.T) {
    tests := []struct {
        set1 []int
        set2 []int
        intersect string
        diff string
        sdiff string
    }{
        {[]int{1, 44, 123, 678}, []int{2, 123, 234, 678, 999},
        "{123 678}", "{1 44}", "{1 2 44 234 999}"},
    }

    for _, test := range tests {
        set1 := &IntSet{[]uint64{}}
        set2 := &IntSet{[]uint64{}}

        set1.AddAll(test.set1...)
        set2.AddAll(test.set2...)

        got := set1.Copy()
        got.IntersectWith(set2)
        if got.String() != test.intersect {
            t.Errorf("intersect failed. want: %v, got: %v", test.intersect, got.String())
        }
        got = set1.Copy()
        got.DifferenceWith(set2)
        if got.String() != test.diff {
            t.Errorf("diff failed. want: %v, got: %v", test.diff, got.String())
        }
        got = set1.Copy()
        got.SymmetricDifference(set2)
        if got.String() != test.sdiff {
            t.Errorf("sdiff failed. want: %v, got: %v", test.sdiff, got.String())
        }
    }
}
