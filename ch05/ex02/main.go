package main

import (
    "fmt"
    "os"
    "golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "countTags: %v\n", err)
        os.Exit(1)
    }
    for tag, cnt := range countTags(make(map[string]int), doc) {
        fmt.Printf("%s: %d\n", tag, cnt)
    }
}

func countTags(tagCnt map[string]int, n *html.Node) map[string]int {
    if n == nil {
        return tagCnt
    } else {
        if n.Type == html.ElementNode {
            tag := n.Data
            tagCnt[tag]++
        }

        tagCnt = countTags(tagCnt, n.FirstChild)
        tagCnt = countTags(tagCnt, n.NextSibling)
        return tagCnt
    }
}
