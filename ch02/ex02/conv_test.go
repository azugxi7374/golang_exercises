package main

import (
    "math"
    // "math/rand"
    "fmt"
    // "strings"
    "testing"
)

func TestString(t *testing.T) {
    var test = []struct {
        input float64
        want string
    }{
        { -123.4, "-123.4°C,-123.4°F,-123.4m,-123.4ft,-123.4lb,-123.4kg" },
    }
    for _, test := range test {
        if got := toStringCat(test.input); got != test.want {
            t.Errorf("String() failed. input: %d, want: %s, got: %s",
            test.input, test.want, got)
        }
    }
}
func toStringCat(v float64) string {
    return fmt.Sprintf("%s,%s,%s,%s,%s,%s",
    Celsius(v),Fahrenheit(v),Meter(v),Feet(v),Pound(v),KGram(v),
)
}


func TestConv(t *testing.T) {
    var tests = []struct {
        input float64
        cf Fahrenheit
        fc Celsius
        mft Feet
        ftm Meter
        lbkg KGram
        kglb Pound
    }{
        {2,
        35.6, -16.6667, 6.56168, 0.6096,  0.907185, 4.40925},
    }
    for _, test := range tests {
        gcf, gfc, gmft, gftm, glbkg, gkglb := run(test.input)
        if !nearlyEquals(float64(test.cf), float64(gcf)) ||
        !nearlyEquals(float64(test.fc), float64(gfc)) ||
        !nearlyEquals(float64(test.mft), float64(gmft)) ||
        !nearlyEquals(float64(test.ftm), float64(gftm)) ||
        !nearlyEquals(float64(test.lbkg), float64(glbkg)) ||
        !nearlyEquals(float64(test.kglb), float64(gkglb)) {
            t.Errorf("input: %f\nwant: %s,%s,%s,%s,%s,%s\ngot: %s,%s,%s,%s,%s,%s",
            test.input,
            test.cf, test.fc, test.mft, test.ftm, test.lbkg, test.kglb,
            gcf, gfc, gmft, gftm, glbkg, gkglb,
        )
    }
}
}

func nearlyEquals(a float64, b float64) bool{
    return math.Abs(a - b) < 1e-2
}

