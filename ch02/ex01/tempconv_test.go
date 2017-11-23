package tempconv

import (
    "math"
    "math/rand"
    "strings"
    "testing"
)

func TestString(t *testing.T) {
    var test = []struct {
        input float64
        want string
    }{
        { 0, "0°C,0°F,0°K" },
        { -123.4, "-123.4°C,-123.4°F,-123.4°K" },
    }
    for _, test := range test {
        if got := toStringCFK(test.input); got != test.want {
            t.Errorf("String() failed. input: %d, want: %s, got: %s",
            test.input, test.want, got)
        }
    }
}

func TestKelvin(t *testing.T){
    var test = []struct {
        input Celsius
        want Kelvin
    }{
        { AbsoluteZeroC, Kelvin(0)},
        { Celsius(-273.15 + 12345), Kelvin(0 + 12345)},
    }
    for _, test := range test {
        if got := CToK(test.input); ! nearlyEquals(float64(got), float64(test.want)) {
            t.Errorf("Kelvin failed. input: %s, want: %s, got: %s",
            test.input, test.want, got)
        }
    }
}

func TestCFKC(t *testing.T) {
    N := 10
    for i:=0; i<N; i++ {
        input := rand.Float64()

        want := Celsius(input)
        got := KToC(FToK(CToF(want)))
        if ! nearlyEquals(float64(want), float64(got)) {
            t.Errorf("CFKC failed. input: %d, want: %s, got: %s",
            input, want.String(), got.String())
        }
    }
}

func TestCKFC(t *testing.T) {
    N := 10
    for i:=0; i<N; i++ {
        input := rand.Float64()

        want := Celsius(input)
        got := FToC(KToF(CToK(want)))
        if ! nearlyEquals(float64(want), float64(got)) {
            t.Errorf("CFKC failed. input: %d, want: %s, got: %s",
            input, want.String(), got.String(),
        )
    }
}
}

func nearlyEquals(a float64, b float64) bool{
    return math.Abs(a - b) < 1e-9
}

func toStringCFK(f float64) string {
    return strings.Join(
        []string{Celsius(f).String(), Fahrenheit(f).String(), Kelvin(f).String()},
        ",",
    )
}


