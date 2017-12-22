package main

import (
    "testing"
)

func TestSha256Diff(t *testing.T){
    var tests = []struct {
        s1 string
        s2 string
        want int
    } {
        {"x", "X", 125},
        {"aaa", "aaa", 0},
    }

    for _, test := range tests {
        got := sha256diff(test.s1, test.s2)
        if got != test.want {
            t.Errorf("s1: %s, s2: %s, want: %d, got: %d", test.s1, test.s2, test.want, got)
        }
    }
}
