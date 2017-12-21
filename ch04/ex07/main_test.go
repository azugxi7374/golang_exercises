package main

import (
	"testing"
)

func TestShaString(t *testing.T) {
	var tests = []struct {
		input string
		want string
	}{
        {"hello ハロー, world 世界！", "！界世 dlrow ,ーロハ olleh"},
        {"", ""},
	}

    for _, test := range tests {
        got := string(reverse([]byte(test.input)))

        if got != test.want {
            t.Errorf("input: [%v], want: [%v], got: [%v]", test.input, test.want, got)
        }
    }
}
