package main

import (
    "bytes"
    "fmt"
    "os"
)

func comma(s string) string {
    var buf bytes.Buffer

    numStartIdx := 0
    for ; !('0' <= s[numStartIdx] && s[numStartIdx]<='9'); numStartIdx++ { }

    intEndIdx := numStartIdx
    for ; intEndIdx < len(s) &&  s[intEndIdx] != '.'; intEndIdx++ {}

    for i:=0; i< len(s); i++ {
        buf.WriteByte(s[i])
        if numStartIdx <= i && i< intEndIdx - 1 && (intEndIdx - 1 - i) % 3 == 0 {
            buf.WriteByte(',')
        }
    }
    return buf.String()
}

func main(){
    fmt.Println(comma(os.Args[1]))
}
