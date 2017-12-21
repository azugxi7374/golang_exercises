package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var flgA = flag.Int("a", 0, "algorightm")

func main() {
	flag.Parse()

	var str string
    if *flgA < 0 || 2 < *flgA {
		os.Exit(1)
	}

    fmt.Print("string <- ")
    fmt.Scan(&str)
    fmt.Println(toShaString(str, *flgA))
}

func toShaString(s string, algo int) string {
    if algo == 0 {
        return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
    } else if algo == 1 {
        return fmt.Sprintf("%x", sha512.Sum384([]byte(s)))
    } else {
        return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
    }
}
