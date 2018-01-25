package main

import (
    "bytes"
    "fmt"
    "testing"
)

func TestLineWordCounter(t *testing.T) {
    tests := []struct {
        input []string
    }{
        {[]string{
            "hoge hoge",
            "aaaaa",
        }},
        {[]string{}},
    }

    for _, test := range tests {
        ln := int64(0)
        joined := ""
        for _, s := range test.input {
            ln += int64(len(s))
            joined += s
        }

        var buf bytes.Buffer
        cw, bytes := CountingWriter(&buf)
        for _, s := range test.input {
            fmt.Fprintf(cw, s)
        }
        if *bytes != ln {
            t.Errorf("byte length wrong. input: %v, want: %v, got: %v",
            test.input, ln, *bytes)
        }
        bufstr := buf.String()
        if bufstr != joined {
            t.Errorf("written string wrong. input: %v, want: %v, got: %v",
            test.input, joined, bufstr)
        }
    }
}
