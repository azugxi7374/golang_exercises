package ftp

import (
    "bufio"
    "fmt"
    "io"
    "net"
    "net/textproto"
    "strings"
    //"time"
)

// var cnt int = 0
func HandleConn(c net.Conn) {
    defer c.Close()

    // fmt.Println("login")
    handleLogin(c)
    // fmt.Println("login...ok")

    for {
        str := readWrite(c)
        if strings.HasPrefix(str, "PORT") {
            io.WriteString(c, "200-_PORT\n")
        } else {
            io.WriteString(c, fmt.Sprintf("ok. %s\n", str))
        }
        // cnt += 1
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

var buf []byte = make([]byte, 1024)
/*
func readWrite(c net.Conn) string {
    c.Read(buf)
    str := string(buf)
    fmt.Print(str)
    return str
}
*/
func readWrite(c net.Conn) string {
    r := textproto.NewReader(bufio.NewReader(c))
    str, _ := r.ReadLine()
    fmt.Println(str)
    return str
}








