package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    fetchCopy(os.Args[1:], os.Stdout)
}

func fetchCopy(urlList []string, dst io.Writer) {
    for _, url := range urlList {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        _, err = io.Copy(dst, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
    }
}

