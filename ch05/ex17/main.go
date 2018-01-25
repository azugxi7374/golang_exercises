package main

import (
    "fmt"
    "net/http"
    "os"

    "golang.org/x/net/html"
)

func main() {
    url := os.Args[1]
    tags := os.Args[2:]
    nodes, _ := findTags(url, tags...)
    for _, n := range nodes {
        fmt.Printf("<%s", n.Data)
        for _, a := range n.Attr {
            fmt.Printf(" %s='%s'", a.Key, a.Val)
        }
        fmt.Println("/>")
    }
}

func findTags(url string, tags ...string) ([]*html.Node, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)
    if err != nil {
        return nil, err
    }

    return ElementsByTagName(doc, tags...), nil
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
    return forEachNode(doc, hasTag(name))
}

func forEachNode(n *html.Node, filter func(*html.Node) bool) []*html.Node {
    nodes := []*html.Node{}

    if filter(n) {
        nodes = append(nodes, n)
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        ret := forEachNode(c, filter)
        nodes = append(nodes, ret...)
    }
    return nodes
}

func hasTag(tags []string) func(*html.Node) bool {
    return func(n *html.Node) bool {
        if n.Type == html.ElementNode {

            for _, tag := range tags {
                if n.Data == tag {
                    return true
                }
            }
        }
        return false
    }
}
