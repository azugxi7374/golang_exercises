package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "testing"
)


// httpが含まれる
func TestAddPrefix(t *testing.T) {
    var tests = []struct {
        input string
        want string
    }{
        {"gopl.io", "http://gopl.io"},
        {"http://gopl.io", "http://gopl.io"},
        {"", "http://"},
    }
    for _, test := range tests {
        if got := addHttp(test.input); got != test.want {
            t.Errorf("input: %s, expected: %s, but was: %s", test.input, test.want, got)
        }
    }
}


func TestFetch(t *testing.T) {
    var tests = [][]string {
        []string {"http://gopl.io", "http://gopl.io"},
    }
    for _, test := range tests {
        var bout = new(bytes.Buffer)
        fetchCopy(test, bout)
        got := bout.String()

        want := ""
        for _, url := range test {
            s := fetchWithStatus(url)
            // fetchWithStatusは200OKで始まる
            if !strings.HasPrefix(s, "200 OK") {
                t.Errorf("dont has prefix 200OK.\n[%s]", s[:100])
            }
            want += s
        }

        // dstにfetchWithStatusの場合と同じ文字列が書き込まれている
        if got != want {
            t.Errorf("not matched. want(len=%d):\n[%s]\ngot(len=%d):\n[%s]\n", len(want), want[:100], len(got), got[:100])
        }
    }
}

func fetchWithStatus(url string) string {
    resp, _ := http.Get(url)
    b, _ := ioutil.ReadAll(resp.Body)
    return fmt.Sprintf("%s\n%s", resp.Status, b)
}

