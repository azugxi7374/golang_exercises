package main

import "fmt"

const (
	SIZE = 5
)

func main() {
	input := [...]int{1, 2, 3, 4, 5}
	out := input
	reverse(&out)
	fmt.Printf("input: %v, output: %v\n", input, out)
}

func reverse(a *[SIZE]int) {
	for i := 0; i < len(a)-1-i; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
}
