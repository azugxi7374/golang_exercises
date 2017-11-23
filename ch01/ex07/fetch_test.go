package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "testing"
)

// dstにresp.Bodyと同じ文字列が書き込まれている
func TestFetch(t *testing.T) {
    var tests = [][]string {
        []string {"http://gopl.io", "http://gopl.io"},
    }
    for _, test := range tests {
        var bout = new(bytes.Buffer)
        fetchCopy(test, bout)
        got := bout.String()

        want := concatRespBody(test)

        if got != want {
            t.Errorf("not matched. want:\n%s\ngot:\n%s\n", want, got)
        }
    }
}

func concatRespBody(urlList []string) string {
    ret := ""
    for _, url := range urlList {
        resp, _ := http.Get(url)
        b, _ := ioutil.ReadAll(resp.Body)
        resp.Body.Close()

        ret += fmt.Sprintf("%s", b)
    }
    return ret
}
