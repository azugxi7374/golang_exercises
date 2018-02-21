package main

import (
    "flag"
    "fmt"
    "log"
    "net"
    "./ftp"
)

var port = flag.Int("port", 8000, "port")

func main() {
    flag.Parse()
    listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
    if err != nil {
        log.Fatal(err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        go ftp.HandleConn(conn)
    }
}
