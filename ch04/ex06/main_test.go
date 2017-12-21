package main

import (
	"testing"
)

func TestShaString(t *testing.T) {
	var tests = []struct {
		input string
		want string
	}{
        {"h e  ll  \t \n  o,  世 界      ", "h e ll o, 世 界 "},
	}

    for _, test := range tests {
        got := string(removeDupSpace([]byte(test.input)))

        if got != test.want {
            t.Errorf("input: [%v], want: [%v], got: [%v]", test.input, test.want, got)
        }
    }
}
