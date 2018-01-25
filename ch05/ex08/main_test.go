package main

import (
    "testing"

    "golang.org/x/net/html"
)

func TestStartEndElement(t *testing.T){
    var tests = []struct {
        url string
        id string
        want string
    }{
        {"https://golang.org", "about", "div"},
        {"https://golang.org", "start", "a"},
        {"https://golang.org", "hogehoge", ""},
    }

    tag := func(node *html.Node) string {
        if node != nil && node.Type == html.ElementNode {
            return node.Data
        }
        return ""
    }

    for _, test := range tests {
        got, _ := findId(test.url, test.id)

        if tag(got) != test.want {
            t.Errorf("Failed. id: %s, want: %s, got: %s", test.id, test.want, tag(got))
        }
    }
}
