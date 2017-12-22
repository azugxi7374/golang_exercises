package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ccnt := wordfreq(os.Args[1])

	for w, n := range ccnt {
		fmt.Printf("%s\t\t%d\n", w, n)
	}
}

func wordfreq(filename string) map[string]int {
	file, _ := os.Open(filename)

	defer file.Close()

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	cnt := make(map[string]int)

	for input.Scan() {
		cnt[input.Text()]++
	}
	return cnt
}
