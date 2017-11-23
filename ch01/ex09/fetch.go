package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    fetchCopy(os.Args[1:], os.Stdout)
}

func fetchCopy(urlList []string, dst io.Writer) {
    for _, url := range urlList {

        resp, err := http.Get(addHttp(url))
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }

        _, err = fmt.Fprintln(dst, resp.Status)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }

        _, err = io.Copy(dst, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
    }
}

func addHttp(url string) string {
    prefix := "http://"
    if strings.HasPrefix(url, prefix) {
        return url
    } else {
        return prefix + url
    }
}
