package main

import (
    "strings"
)

func join(sep string, vals ...string) string {
    return strings.Join(vals, sep)
}

