package main

import (
    "testing"
)

func TestShaString(t *testing.T){
    var tests = []struct {
        s string
        algo int
        want string
    } {
        {"x", 0, "2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881"},
        {"x", 1, "d752c2c51fba0e29aa190570a9d4253e44077a058d3297fa3a5630d5bd012622f97c28acaed313b5c83bb990caa7da85"},
        {"x", 2, "a4abd4448c49562d828115d13a1fccea927f52b4d5459297f8b43e42da89238bc13626e43dcb38ddb082488927ec904fb42057443983e88585179d50551afe62"},
    }

    for _, test := range tests {
        got := toShaString(test.s, test.algo)
        if got != test.want {
            t.Errorf("s: %s, algo: %d, want: %s, got: %s", test.s, test.algo, test.want, got)
        }
    }
}
