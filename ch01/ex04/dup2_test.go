package main

import (
    "testing"
    // "strings"
)

func TestDup2(t *testing.T) {
    var input = []string { "file1.txt", "file2.txt", "file3.txt" }
    var want = []string {
        "2\t[file1.txt]\tabcabcabc",
        "5\t[file2.txt, file3.txt]\thoge",
        "3\t[file1.txt, file2.txt, file3.txt]\tssttssttsstt",
    }

    if got := run(input); !sliceEquals(got, want) {
        t.Errorf("\ninput: %v\nexpected: %v\nbut was: %v",
            input, want, got)
    }
}

func sliceEquals(a []string, b []string) bool {
    if len(a) != len(b) {
        return false
    } else {
        for i, _ := range a {
            if a[i] != b[i] {
                return false
            }
        }
        return true
    }
}
