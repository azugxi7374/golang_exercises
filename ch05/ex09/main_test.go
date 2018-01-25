package main

import (
    "testing"
)

func TestStartEndElement(t *testing.T){
    var tests = []struct {
        input string
        want string
    }{
        {"abc $def $xxx e", "abc defdef xxxxxx e"},
        {"abc def xxx e", "abc def xxx e"},
        {"", ""},
    }

    for _, test := range tests {
        got := expand(test.input, twice)

        if got != test.want {
            t.Errorf("Failed. input: %s, want: %s, got: %s", test.input, test.want, got)
        }
    }
}
