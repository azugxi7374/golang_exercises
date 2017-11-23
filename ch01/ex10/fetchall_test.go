package main

import (
    "bytes"
    "net/http"
    "fmt"
    "io"
    "io/ioutil"
    "testing"
)

// fetchallで書き込まれるものとhttp.Getが一致
func TestFetchAll(t *testing.T) {
    var tests = []struct{
        url string
    }{
        { "https://ja.wikipedia.org/wiki/Go_(%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9E)"},
    }

    for _, test := range tests {
        out := new(bytes.Buffer)
        fetchall([]string{test.url}, []io.Writer{out})
        got1 := out.String()

        resp, _ := http.Get(test.url)
        b, _ := ioutil.ReadAll(resp.Body)
        got2 := fmt.Sprintf("%s", b)
        if got1 != got2 {
            t.Errorf("failed.\ngot1:%s\ngot2:%s\n", got1[:100], got2[:100])
        }
    }
}
