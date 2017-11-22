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
        {[]string {"hello", "world"}, "0: hello\n1: world\n"},
        {[]string {}, ""},
    }

    for _, test := range tests {
        if got := echo(test.input); got != test.want {
            t.Errorf("echo(" + strings.Join(test.input, ",") +
            ") expected:" + test.want + " but was: " + got)
        }
    }
}
