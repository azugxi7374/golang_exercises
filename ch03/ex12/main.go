package main

import (
    "fmt"
    "os"
)

func anagram(s1 string, s2 string) bool {
    r1 := []rune(s1)
    r2 := []rune(s2)
    if len(r1) != len(r2) {
        return false
    } else {
        n := len(r1)
        done := make([]bool, n)
        for i:=0; i<n; i++ {
            found := false
            for jj:=0; !found && jj<n; jj++ {
                if !done[jj] && r1[i] == r2[jj] {
                    found = true
                    done[jj] = true
                }
            }
            if !found {
                return false
            }
        }
        return true
    }
}

func main(){
    fmt.Println(anagram(os.Args[1], os.Args[2]))
}
