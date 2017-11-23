package popcount

import (
    "testing"
    "math/rand"
)

func rand64() uint64 {
    return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}

// PopCount0, PopCount1 が正しい
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
        if got := PopCount1(test.input); got != test.want {
            t.Errorf("PopCount1 failed. input:%d, want:%d, got:%d", test.input, test.want, got)
        }
    }
}


const B_MAX = 65536

func BenchmarkPopCount0(b *testing.B){
    for t:=0; t<b.N; t++ {
        PopCount0(rand64())
    }
}
func BenchmarkPopCount1(b *testing.B){
    for t:=0; t<b.N; t++ {
        PopCount1(rand64())
    }
}
