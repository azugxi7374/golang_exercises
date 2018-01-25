package main

import (
    "testing"
)

func TestMaxMin(t *testing.T){
    var tests = []struct {
        input []string
        sep string
        want string
    }{
        {[]string{"a", "b", "ccc"}, ",", "a,b,ccc"},
        {[]string{"a", "b", "ccc"}, "___", "a___b___ccc"},
        {[]string{}, ",", ""},
    }

    for _, test := range tests {
        got := join(test.sep, test.input...)
        if test.want != got {
            t.Errorf("Failed. input: %v, sep: %v, want: %v, got: %v",
            test.input, test.sep, test.want, got)
        }
    }
}
