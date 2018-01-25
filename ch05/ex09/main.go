package main

import (
    "fmt"
    "regexp"
)

var extraction = regexp.MustCompile(`(\$\w*)`)

func expand(s string, f func(string) string) string {
    repStr := func(s1 string) string {
        return f(s1[1:])
    }
    return extraction.ReplaceAllStringFunc(s, repStr)
}

func twice(s string) string {
    return fmt.Sprintf("%s%s", s, s)
}
