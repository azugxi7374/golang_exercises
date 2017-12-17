package main

import (
    "testing"
)

func TestCommma310(t *testing.T) {
    var tests = []struct {
        input string
    }{
        {"1"},
        {"12"},
        {"123"},
        {"1234"},
        {"12345678901234567890"},
    }

    for _, test := range tests {
        want := commaOriginal(test.input)
        got := comma310(test.input)
        if got != want {
            t.Errorf("Failed. input:%s, want:%s, got:%s", test.input, want, got)
        }
    }
}

