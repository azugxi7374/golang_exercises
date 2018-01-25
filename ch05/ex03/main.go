package main

import (
    "fmt"
    "os"
    "strings"

    "golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }
    for _, str := range visitTextNode(doc, nil) {
        fmt.Printf("%s\n", str)
    }
}

func visitTextNode(n *html.Node, strList []string) []string {
    if n == nil {
        return strList
    } else if n.Type != html.TextNode && (n.Data == "script" || n.Data == "style") {
        // skip!
        return strList
    }
    if n.Type == html.TextNode {
        str := strings.TrimSpace(n.Data)
        if len(str) > 0 {
            strList = append(strList, n.Data)
        }
    }

    strList = visitTextNode(n.FirstChild, strList)
    strList = visitTextNode(n.NextSibling, strList)
    return strList
}
