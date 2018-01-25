package main

import (
    "fmt"
    "testing"
)

func TestLineWordCounter(t *testing.T) {
    tests := []struct {
        input string
        wcnt  int
        lcnt  int
    }{
        {
            "hoge hoge fuga f ug a\npiyo pi! yo!!\npiyopiyo",
            10,
            3,
        },
        {
            "hello world!",
            2,
            1,
        },
    }

    for _, test := range tests {
        wc := WordCounter(0)
        lc := LineCounter(0)

        fmt.Fprintf(&wc, test.input)
        if int(wc) != test.wcnt {
            t.Errorf("wordcounter failed. input: %v, want: %v, got: %v",
            test.input, test.wcnt, wc)
        }
        fmt.Fprintf(&lc, test.input)
        if int(lc) != test.lcnt {
            t.Errorf("linecounter failed. input: %v, want: %v, got: %v",
            test.input, test.lcnt, lc)
        }
    }
}
