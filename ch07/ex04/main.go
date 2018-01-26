package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"

    "golang.org/x/net/html"
)
func main() {
    contents := ""
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        contents += sc.Text()
        contents += "\n"
    }

    str, _ := outline(contents)
    fmt.Println(str)
}

type StringReader struct{
    contents string
    next int
}
func (sr *StringReader)Read(p []byte) (int, error) {
    if sr.next >= len(sr.contents) {
        return 0, io.EOF
    }
    n := len(sr.contents) - sr.next
    if n > len(p) {
        n = len(p)
    }

    copy(p, []byte(sr.contents[sr.next: sr.next + n]))

    sr.next += n
    return n, nil
}
func newReader(contents string) io.Reader {
    return &StringReader{contents, 0}
}

func outline(docStr string) (string, error) {
    doc, err := html.Parse(newReader(docStr))
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
