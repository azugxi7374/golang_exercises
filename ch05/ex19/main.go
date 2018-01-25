package main

import "fmt"

func main() {
    fmt.Println(hoge())
}

func hoge() (ret string) {
    defer func() {
        if p := recover(); p != nil {
            ret = p.(string)
        }
    }()

    panic("hoge")
}
