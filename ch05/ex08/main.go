package main

import (
    "fmt"
    "net/http"
    "os"

    "golang.org/x/net/html"
)

func main() {
    url := os.Args[1]
    id := os.Args[2]
    n, _ := findId(url, id)
    if n != nil {
        fmt.Printf("<%s", n.Data)
        for _, a := range n.Attr {
            fmt.Printf(" %s='%s'", a.Key, a.Val)
        }
        fmt.Println("/>")
    } else {
        fmt.Println("Not Found.")
    }
}

func findId(url string, id string) (*html.Node, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)
    if err != nil {
        return nil, err
    }

    return ElementById(doc, id), nil
}

func ElementById(doc *html.Node, id string) *html.Node {
    return forEachNode(doc, hasId(id), nil)
}

func forEachNode(n *html.Node, pre, post func(*html.Node) bool) *html.Node {
    if pre != nil {
        fin := pre(n)
        if fin {
            return n
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        ret := forEachNode(c, pre, post)
        if ret != nil {
            return ret
        }
    }

    if post != nil {
        fin := post(n)
        if fin {
            return n
        }
    }

    return nil
}

func hasId(id string) func(*html.Node) bool {
    return func(n *html.Node) bool {
        if n.Type == html.ElementNode {
            for _, a := range n.Attr {
                if a.Key == "id" && a.Val == id {
                    return true
                }
            }
        }
        return false
    }
}
