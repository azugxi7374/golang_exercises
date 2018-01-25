package main

import (
    "strings"
    "testing"

    "golang.org/x/net/html"
)

func TestStartEndElement(t *testing.T){
    var tests = []struct {
        url string
    }{
        {"https://golang.org"},
    }

    for _, test := range tests {
        doc, _ := outline(test.url)

        _, err := html.Parse(strings.NewReader(doc))
        if err != nil {
            t.Errorf("Failed. %v", err)
        }
    }
}
