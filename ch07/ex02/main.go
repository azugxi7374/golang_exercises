package main

import (
    "io"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
    cw := CWriter{0, w}
    return &cw, &(cw.cnt)
}

type CWriter struct {
    cnt int64
    writer io.Writer
}

func (cw *CWriter) Write(b []byte) (int, error) {
    n, err := cw.writer.Write(b)
    cw.cnt += int64(n)
    return n, err
}
