package main

import (
    "testing"
    "strings"
)

func BenchmarkJoinPlus(b *testing.B){
    arg := []string {"a", "b", "c", "d", "e"}
    for i:= 0; i<b.N; i++ {
        JoinPlus(arg)
    }
}

func BenchmarkJoinStrings(b *testing.B){
    arg := []string {"a", "b", "c", "d", "e"}
    for i:= 0; i<b.N; i++ {
        JoinStrings(arg)
    }
}


func Test(t *testing.T){
    arg := []string {"a", "b", "c", "d", "e"}
    jp := JoinPlus(arg)
    js := JoinStrings(arg)

    if jp != js {
        t.Error(jp + " != " + js)
    }
}


func JoinPlus(args []string) string{
    s, sep := "", ""
    for _, arg := range args {
        s += sep + arg
        sep = ","
    }
    return s
}
func JoinStrings(args []string) string{
    return strings.Join(args, ",")
}


