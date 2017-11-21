package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    fmt.Print(echo(os.Args[1:]))
}

func echo(args []string) string{
    ret := ""
    for i, s:= range args {
        ret += strconv.Itoa(i) + ": " + s + "\n"
    }
    return ret
}

