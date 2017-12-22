package main

import (
	"testing"
)

func TestShaString(t *testing.T) {
	var tests = []struct {
		input []string
		want []string
	}{
		{[]string{}, []string{}},
		{[]string{"a", "a", "a"}, []string{"a"}},
        {[]string{"aaa", "bbb", "bbb", "bbb", "ccc", "bbb", "bbb"}, []string{"aaa", "bbb", "ccc", "bbb"}},
	}

    for _, test := range tests {
        got := make([]string, len(test.input))
        copy(got, test.input)
        got = removeDup(got)

        if !sliceEquals(test.want, got) {
            t.Errorf("input: %v, want: %v, got: %v", test.input, test.want, got)
        }
    }
}

func sliceEquals(a []string, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, _ := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
