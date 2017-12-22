package main

import "fmt"

func main() {
	input := []string{"aaa", "bbb", "bbb", "bbb", "ccc"}
    out := make([]string, len(input))
    copy(out, input)
    out = removeDup(out)
    fmt.Println(input, " -> ", out)
}

func removeDup(a []string) []string{
    for i:= 1; i < len(a); {
        if a[i] == a[i-1] {
            copy(a[i:], a[i+1:])
            a = a[:len(a)-1]
        }else{
            i++
        }
    }
    return a
}
