package main

import (
    "fmt"
    "sort"
)

var prereqs = map[string][]string{
    "algorithms": {"data structures"},
    "calculus":   {"linear algebra"},

    "compilers": {
        "data structures",
        "formal languages",
        "computer organization",
    },

    "data structures":       {"discrete math"},
    "databases":             {"data structures"},
    "discrete math":         {"intro to programming"},
    "formal languages":      {"discrete math"},
    "networks":              {"operating systems"},
    "operating systems":     {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
}

func main() {
    sorted, ok := topoSort(prereqs)
    if ok {
        for i, course := range sorted {
            fmt.Printf("%d:\t%s\n", i+1, course)
        }
    }
}

func topoSort(m map[string][]string) ([]string, bool) {
    var order []string
    seen := make(map[string]bool)
    done := make(map[string]bool)

    loop := false
    var visitAll func(items []string)

    visitAll = func(items []string) {
        for _, item := range items {
            if loop || seen[item] && !done[item] {
                loop = true
                return
            }
            if !seen[item] {
                seen[item] = true
                visitAll(m[item])
                order = append(order, item)
                done[item] = true
            }
        }
    }

    var keys []string
    for key := range m {
        keys = append(keys, key)
    }

    sort.Strings(keys)
    visitAll(keys)
    if loop {
        return nil, false
    } else {
        return order, true
    }
}
