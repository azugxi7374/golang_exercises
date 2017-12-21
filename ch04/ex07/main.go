package main

import (
	"fmt"
)

func main() {
	input := "hello ハロー, world 世界！"
	b := []byte(input)
	b = reverse(b)
	out := string(b)
	fmt.Printf("[%s] -> [%s]\n", input, out)
}

func reverse(ustr []byte) []byte {
	for s1 := 0; s1 < len(ustr); {
		s2 := s1 + 1
		for s2 < len(ustr) && ustr[s2]>>6 == 2 {
			s2++
		}
		for i := 0; s1+i < s2-1-i; i++ {
			ustr[s1+i], ustr[s2-1-i] = ustr[s2-1-i], ustr[s1+i]
		}
		s1 = s2
	}
	for i := 0; i < len(ustr)-1-i; i++ {
		ustr[i], ustr[len(ustr)-1-i] = ustr[len(ustr)-1-i], ustr[i]
	}
	return ustr
}
