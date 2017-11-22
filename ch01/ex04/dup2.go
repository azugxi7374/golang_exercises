package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

func main() {
    files := os.Args[1:]

    for _, line := range run(files) {
        fmt.Println(line)
    }
}

func run(files []string) []string {
	counts := make(map[string]int)
    incfiles := make(map[string]map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, counts, incfiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			} else {
				countLines(f, counts, incfiles)
				f.Close()
			}
		}
	}
    sorted := []string{}
    for line, _ := range counts {
        sorted = append(sorted, line)
    }
    sort.Strings(sorted)

    res := []string{}
	for _, line := range sorted {
        n := counts[line]
		if n > 1 {
            var fa []string
            for f, _ := range incfiles[line] {
                fa = append(fa, f)
            }
            sort.Strings(fa)
            str := fmt.Sprintf("%d\t[%s]\t%s", n, strings.Join(fa, ", "), line)
            res = append(res, str)
		}
	}
    return res
}

func countLines(f *os.File, counts map[string]int, incfiles map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
        t := input.Text()
		counts[t]++
        if incfiles[t] == nil {
            incfiles[t] = make(map[string]int)
        }
        incfiles[t][f.Name()]++
	}
}

