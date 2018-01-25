package main

import (
    "testing"
)

func TestStartEndElement(t *testing.T){
    var tests = []struct {
        input map[string][]string
        want bool
    }{
        {prereqs, true},
        {make(map[string][]string), true},
        {map[string][]string{
            "calculus":   {"linear algebra"},
            "linear algebra": {"calculus"},
        }, false},
    }

    for _, test := range tests {
        _, got := topoSort(test.input)

        if got != test.want {
            t.Errorf("Failed. input: %v, want: %v, got: %v", test.input, test.want, got)
        }
    }
}
