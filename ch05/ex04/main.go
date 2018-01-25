package main

import (
    "fmt"
    "os"
    "golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
        os.Exit(1)
    }
    for _, link := range visit(nil, doc) {
        fmt.Println(link)
    }
}

var LINK_TAGS = [][]string{{"a", "href"}, {"img", "src"}, {"link", "href"}, {"script", "src"}}

func visit(links []string, n *html.Node) []string {
    if n == nil {
        return links
    } else {
        if n.Type == html.ElementNode {
            key := ""
            for _, elem := range LINK_TAGS {
                if elem[0] == n.Data {
                    key = elem[1]
                }
            }
            if key != "" {
                for _, a := range n.Attr {
                    if a.Key == key {
                        str := fmt.Sprintf("%s: %s", a.Key, a.Val)
                        links = append(links, str)
                    }
                }
            }
        }

        links = visit(links, n.FirstChild)
        links = visit(links, n.NextSibling)
        return links
    }
}
