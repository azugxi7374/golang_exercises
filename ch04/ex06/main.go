package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := "h e  ll  \t \n  o,  世 界      "
	b := []byte(input)
	b = removeDupSpace(b)
    out := string(b)
	fmt.Printf("[%s] -> [%s]\n", input, out)
}

func removeDupSpace(ustr []byte) []byte {
	cur := 0
	prevIsSpace := false
	for cur < len(ustr) {
		next := cur + 1
		for next < len(ustr) && (ustr[next] >> 6) == 2 {
			next++
		}
		rn := 0
		for i := cur; i < next; i++ {
			rn = (rn << 8) | int(ustr[i])
		}
		curIsSpace := unicode.IsSpace(rune(rn))

		if curIsSpace {
			if !prevIsSpace {
				ustr[cur] = 0x20
				cur++
			}
			copy(ustr[cur:], ustr[next:])
			ustr = ustr[:len(ustr)-(next-cur)]
		} else {
			cur = next
		}
		prevIsSpace = curIsSpace
	}
	return ustr
}
