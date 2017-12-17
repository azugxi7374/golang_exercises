package main

import (
    "bytes"
    "fmt"
    "os"
)

func commaOriginal(s string) string {
    n := len(s)
    if n<= 3 {
        return s
    }
    return commaOriginal(s[:n-3]) + "," + s[n-3:]
}

func comma310(s string) string {
    var buf bytes.Buffer
    n:=len(s)
    for i:=0; i< n; i++ {
        buf.WriteByte(s[i])
        if (n - 1 - i) % 3 == 0 && i< n -1 {
            buf.WriteByte(',')
        }
    }
    return buf.String()
}


func main(){
    fmt.Println(comma310(os.Args[1]))
}
