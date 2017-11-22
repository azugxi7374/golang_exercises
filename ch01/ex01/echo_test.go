package main

import (
    "testing"
    "strings"
)

func TestEcho(t *testing.T) {
    var tests = []struct {
        input []string
        want string
    }{
        {[]string {"a", "b", "c"}, "a b c"},
        {[]string {"a"}, "a"},
        {[]string {}, ""},
    }

    for _, test := range tests {
        if got := echo(test.input); got != test.want {
            t.Errorf("echo(" + strings.Join(test.input, ",") + ") = " + got)
        }
    }
}
