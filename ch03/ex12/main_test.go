package main

import (
    "testing"
)

func TestCommma310(t *testing.T) {
    var tests = []struct {
        s1 string
        s2 string
        want bool
    }{
        {"12345", "31542", true},
        {"abcabcab", "ccaaabbb", true},
        {"", "", true},
        {"12345", "1234", false},
        {"abcab", "bbacc", false},
    }

    for _, test := range tests {
        got := anagram(test.s1, test.s2)
        if got != test.want {
            t.Errorf("Failed. input:%s and %s, want:%t, got:%t", test.s1, test.s2, test.want, got)
        }
    }
}

