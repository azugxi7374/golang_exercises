package main

import (
    "io"
)

func LimitReader(r io.Reader, n int64) io.Reader {
    return &LReader{r, int(n)}
}

type LReader struct {
    origin    io.Reader
    remaining int
}

func (lr *LReader) Read(p []byte) (int, error) {
    if lr.remaining == 0 {
        return 0, io.EOF
    } else {
        n, err := lr.origin.Read(p)
        if n > lr.remaining {
            n = lr.remaining
        }
        lr.remaining -= n
        return n, err
    }
}
