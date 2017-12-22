package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5}
    out := make([]int, 5)
    copy(out, input)
    rotate(out, 2)
    fmt.Println(input, " -> rotate 2 ->", out)
}

func rotate(a []int, r int) {
    L := len(a)
    if L == 0 {
        return
    }

    r = (r % L + L) % L

    start := 0
    moved := 0
    for moved < L {
        for to := (start + r) % L; to!= start; to = (to + r) % L {
            a[start], a[to] = a[to], a[start]
            moved++
        }
        moved++

        start++
    }
}
