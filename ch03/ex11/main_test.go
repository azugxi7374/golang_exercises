package main

import (
    "testing"
)

func TestCommma310(t *testing.T) {
    var tests = []struct {
        input string
        want string
    }{
        {"1", "1"},
        {"+12", "+12"},
        {"-1234", "-1,234"},
        {"+123456789.123456789", "+123,456,789.123456789"},
        {"123456789.123456789", "123,456,789.123456789"},
    }

    for _, test := range tests {
        got := comma(test.input)
        if got != test.want {
            t.Errorf("Failed. input:%s, want:%s, got:%s", test.input, test.want, got)
        }
    }
}

