package main

import (
    "fmt"
    "io"
    "bufio"
    "log"
    "net"
    "os"
    "strings"
)

func main() {
    names := make([]string, len(os.Args)-1)
    // cnn := make([]*net.Conn, len(os.Args)-1)
    rd := make([]*bufio.Reader, len(os.Args)-1)
    for i, arg := range os.Args[1:] {
        name, url := argParse(arg)
        conn, err := net.Dial("tcp", url)
        names[i] = name
        if err != nil {
            log.Fatal(err)
        }
        defer conn.Close()
        // mustCopy(os.Stdout, conn)
        // cnn[i] = bufio.NewScanner(conn)
        rd[i] = bufio.NewReader(conn)
    }
    fmt.Println(strings.Join(names, " | "))
    for {
        times := []string{}
        for _, r := range rd {
            t, _ := r.ReadString('\n')
            times = append(times, t[:len(t)-1])
        }
        fmt.Println(strings.Join(times, " | "))
    }
}

func argParse(arg string) (string, string) {
    idx := strings.Index(arg, "=")
    return arg[0:idx], arg[idx+1:]
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}

//!-
