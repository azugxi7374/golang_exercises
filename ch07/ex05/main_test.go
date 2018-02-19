package main

import(
    "io/ioutil"
    "strings"
    "testing"
)

func TestLimitReader(t *testing.T){
    tests := []struct{
        text string
        n int
        want string
    }{
        {"", 999, ""},
        {"abcd", 2, "ab"},
        {"abcd", 555, "abcd"},
        {"aaaaaaaaaaaaaaaaaaa", 0, ""},
    }

    for _, test := range tests {
        lr := & LReader{ strings.NewReader(test.text), test.n }
        gotb, _ := ioutil.ReadAll(lr)
        got := string(gotb)
        if test.want != got {
            t.Errorf("text: %v, n: %v, want: %v, got: %v",
                test.text, test.n, test.want, got)
        }
    }
}





