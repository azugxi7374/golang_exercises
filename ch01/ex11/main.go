package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    var urls []string
    var writers []io.Writer
    i := 1
    dump := len(os.Args) > 1 && os.Args[1] == "--dump"
    if dump { i++ }
    for ; i < len(os.Args); {
        urls = append(urls, os.Args[i])
        i++
        if dump {
            f, _ := os.Create(os.Args[i])
            writers = append(writers, f)
            i++
        } else {
            writers = append(writers, ioutil.Discard)
        }
    }
    start := time.Now()

    fetchall(urls, writers)

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchall(urls []string, writers []io.Writer) {
    ch := make(chan string)
    for i, _ := range urls{
        go fetch(urls[i], ch, writers[i])
    }
    for range urls {
        fmt.Println(<-ch)
    }
}

func fetch(url string, ch chan<- string, out io.Writer) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    nbytes, err := io.Copy(out, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
