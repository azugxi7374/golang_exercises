package main

import (
    "bufio"
    "bytes"
)

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
    b := len(p)

    sc := bufio.NewScanner(bytes.NewReader(p))

    for sc.Scan() {
        _ = sc.Text()
        *lc += 1
    }

    return b, nil
}

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
    b := len(p)

    for len(p) > 0 {
        next, tk, _ := bufio.ScanWords(p, true)

        if tk != nil {
            *wc += 1
        }
        p = p[next:]
    }
    return b, nil
}
