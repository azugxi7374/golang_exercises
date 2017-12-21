package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	ccnt := charcount()
	fmt.Printf("%v\n", ccnt)
}

var categories = []struct {
	name string
	f    func(rune) bool
}{
	{"Control", unicode.IsControl},
	{"Digit", unicode.IsDigit},
	{"Graphic", unicode.IsGraphic},
	{"Letter", unicode.IsLetter},
	{"Lower", unicode.IsLower},
	{"Mark", unicode.IsMark},
	{"Number", unicode.IsNumber},
	{"Print", unicode.IsPrint},
	{"Punctuation", unicode.IsPunct},
	{"Space", unicode.IsSpace},
	{"Symbol", unicode.IsSymbol},
	{"Title", unicode.IsTitle},
	{"Upper", unicode.IsUpper},
}

func charcount() map[string]int {
	cnt := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		for _, ctg := range categories {
			if ctg.f(r) {
				cnt[ctg.name]++
				break
			}
		}
	}
	return cnt
}
