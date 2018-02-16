package treesortstring

import (
    "fmt"
)

type tree struct {
    value       int
    left, right *tree
}

func (t *tree)String() string {
    if t == nil {
        return ""
    } else if t.left == nil && t.right == nil {
        return fmt.Sprintf("%v", t.value)
    } else if t.right == nil {
        return fmt.Sprintf("%v, %v", t.left.String(), t.value)
    } else if t.left == nil {
        return fmt.Sprintf("%v, %v", t.value, t.right.String())
    } else {
        return fmt.Sprintf("%v, %v, %v", t.left.String(), t.value, t.right.String())
    }
}

func add(t *tree, value int) *tree {
    if t == nil {
        t = new(tree)
        t.value = value
        return t
    }
    if value < t.value {
        t.left = add(t.left, value)
    } else {
        t.right = add(t.right, value)
    }
    return t
}

