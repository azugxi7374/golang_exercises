package ftp

import (
    "fmt"
    "io"
    "net"
    //"time"
)

func HandleConn(c net.Conn) {
    defer c.Close()

    fmt.Println("login")
    handleLogin(c)
    fmt.Println("login...ok")

    for {
        readWrite(c)
        io.WriteString(c, "ok\n")
    }
    /*
    for {
        buf := make([]byte, 1024)
        c.Read(buf)
        // for sc.Scan() {
        // t += sc.Text()
        // }
        fmt.Println(string(buf))
        _, err := io.WriteString(c, "aaa")
        if err != nil {
            return // e.g., client disconnected
        }
        time.Sleep(1 * time.Second)
    }        */

}

func handleLogin(c net.Conn) {
    io.WriteString(c, "220\n")
    readWrite(c)
    io.WriteString(c, "331\n")
    readWrite(c)
    io.WriteString(c, "230\n")
    io.WriteString(c, "230\n")
}

func readWrite(c net.Conn) {
    buf := make([]byte, 1024)
    c.Read(buf)
    fmt.Println(string(buf))
}
