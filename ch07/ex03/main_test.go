package treesortstring

import (
    "testing"
)

func TestTreeString(t *testing.T) {
    tests := []struct{
        values []int
        want string
    }{
        {[]int{}, ""},
        {[]int{1}, "1"},
        {[]int{5, 1}, "1, 5"},
        {[]int{2,4,6,3,1,-3,9999,123}, "-3, 1, 2, 3, 4, 6, 123, 9999"},
    }
    for _, test := range tests {
        var root *tree
        for _, v := range test.values {
            root = add(root, v)
        }
        got := root.String()
        if got != test.want {
            t.Errorf("error. input: %v, want %s, got: %s", test.values, test.want, got)
        }
    }
}
