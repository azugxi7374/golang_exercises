package main

import (
    "strings"
    "testing"
)

// NaNやINFをskipしている
func TestCreateSVG(t *testing.T){
    var tests = []struct {
        input func(float64, float64) float64
    }{
        {fTan},
    }

    existsNanInf := func(str string) bool {
        return strings.Contains(str, "NaN") || strings.Contains(str, "Inf")
    }
    for _, test := range tests {
        raw := createSVG(test.input, false)
        skipped := createSVG(test.input, true)

        if !existsNanInf(raw) {
            t.Errorf("raw don't contains NaN nor Inf\n%s")
        }
        if existsNanInf(skipped) {
            t.Errorf("skipped contains NaN or Inf")
        }
    }
}
