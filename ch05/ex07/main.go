package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"

    "golang.org/x/net/html"
)

func main() {
    for _, url := range os.Args[1:] {
        str, _ := outline(url)
        fmt.Println(str)
    }
}

func outline(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)
    if err != nil {
        return "", err
    }

    return forEachNode(doc, startElement, endElement, 0), nil
}

func forEachNode(n *html.Node, pre, post func(*html.Node, int, bool) string, depth int) string {
    ret := ""
    nextDepth := depth
    if pre != nil {
        str := pre(n, depth, n.FirstChild == nil)
        ret += str
        if len(str) > 0 {
            nextDepth += 1
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        ret += forEachNode(c, pre, post, nextDepth)
    }

    if post != nil {
        str := post(n, depth, n.FirstChild == nil)
        ret += str
    }

    return ret
}

func startElement(n *html.Node, depth int, noChildren bool) string {
    str := ""
    switch n.Type {
    case html.ElementNode:
        str = "<"
        str += n.Data

        for _, a := range n.Attr {
            str += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
        }

        if noChildren {
            str += "/>"
        } else {
            str += ">"
        }
    case html.CommentNode:
        str = n.Data
    case html.TextNode:
        str = strings.TrimSpace(n.Data)
    default:
        // NOP
    }
    if len(str) > 0 {
        return fmt.Sprintf("%*s%s\n", depth*2, "", str)
    } else {
        return ""
    }
}

func endElement(n *html.Node, depth int, noChildren bool) string {
    switch n.Type {
    case html.ElementNode:
        if(!noChildren) {
            return fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
        }
    }
    return ""
}
