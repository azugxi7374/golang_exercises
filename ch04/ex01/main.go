package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	fmt.Println(sha256diff(os.Args[1], os.Args[2]))
}

func sha256diff(s1 string, s2 string) int {
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))

	cnt := 0
	for i := 0; i < len(c1); i++ {
		cnt += bytePopCount(c1[i] ^ c2[i])
	}
	return cnt
}

func bytePopCount(b byte) int {
	cnt := 0
	for b > 0 {
		cnt += int(b & 1)
		b = b >> 1
	}
	return cnt
}
