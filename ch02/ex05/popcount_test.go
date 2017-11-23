package popcount

import (
    "testing"
    "math/rand"
)

func rand64() uint64 {
    return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}

// PopCount0, PopCountEx3, PopCountEx4, PopCountEx5 が正しい
func TestPopCount(t *testing.T){
    var tests = []struct {
        input uint64
        want int
    }{
        {0, 0},
        {1, 1},
        {7, 3},
        {5, 2},
        {18446744073709551615, 64},
    }

    for _, test := range tests {
        if got := PopCount0(test.input); got != test.want {
            t.Errorf("PopCount0 failed. input:%d, want:%d, got:%d", test.input, test.want, got)
        }
        if got := PopCountEx3(test.input); got != test.want {
            t.Errorf("PopCountEx3 failed. input:%d, want:%d, got:%d", test.input, test.want, got)
        }
        if got := PopCountEx4(test.input); got != test.want {
            t.Errorf("PopCountEx4 failed. input:%d, want:%d, got:%d", test.input, test.want, got)
        }
        if got := PopCountEx5(test.input); got != test.want {
            t.Errorf("PopCountEx5 failed. input:%d, want:%d, got:%d", test.input, test.want, got)
        }
    }

}

func BenchmarkPopCount0(b *testing.B){
    for t:=0; t<b.N; t++ {
        PopCount0(rand64())
    }
}
func BenchmarkPopCountEx3(b *testing.B){
    for t:=0; t<b.N; t++ {
        PopCountEx3(rand64())
    }
}
func BenchmarkPopCountEx4(b *testing.B){
    for t:=0; t<b.N; t++ {
        PopCountEx4(rand64())
    }
}
func BenchmarkPopCountEx5(b *testing.B){
    for t:=0; t<b.N; t++ {
        PopCountEx5(rand64())
    }
}

