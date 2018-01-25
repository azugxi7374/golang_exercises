package intset

import (
    "bytes"
    "fmt"
)

type IntSet struct {
    words []uint64
}

func (s *IntSet) Has(x int) bool {
    word, bit := x/64, uint(x%64)
    return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
    word, bit := x/64, uint(x%64)
    for word >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(x ...int){
    for _, xx := range x {
        s.Add(xx)
    }
}

func (s *IntSet) Len() int {
    cnt := 0
    for _, w := range s.words {
        for i:=uint(0); w >> i != 0; i++ {
            cnt += int((w>>i) & 1)
        }
    }
    return cnt
}
func (s *IntSet) Remove(x int) {
    word, bit := x/64, uint(x%64)
    if word < len(s.words) {
        s.words[word] &^= 1<<bit
    }
}
func (s *IntSet) Clear() {
    s.words = []uint64{}
}
func (s *IntSet) Copy() *IntSet{
    wds := []uint64{}
    for _, w := range s.words {
        wds = append(wds, w)
    }
    return &IntSet{wds}
}

func (s *IntSet) UnionWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] |= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}
func (s *IntSet) IntersectWith(t *IntSet) {
    for i, _ := range s.words {
        if i < len(t.words) {
            s.words[i] &= t.words[i]
        } else {
            s.words[i] = 0
        }
    }
}
func (s *IntSet) DifferenceWith(t *IntSet) {
    for i, _ := range s.words {
        if i < len(t.words) {
            s.words[i] &^= t.words[i]
        }
    }
}
func (s *IntSet) SymmetricDifference(t *IntSet) {
    for i, _ := range t.words {
        if i < len(s.words) {
            s.words[i] ^= t.words[i]
        }else{
            s.words = append(s.words, t.words[i])
        }
    }
}

func (s *IntSet) Elems() []int{
    ret := []int{}
    for i, w := range s.words {
        for b:=uint64(0); (w>>b) != 0; b++ {
            if (w>>b&1) == 1 {
                ret = append(ret, int(i*64+int(b)))
            }
        }
    }
    return ret
}

func (s *IntSet) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for i, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < 64; j++ {
            if word&(1<<uint(j)) != 0 {
                if buf.Len() > len("{") {
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", 64*i+j)
            }
        }
    }
    buf.WriteByte('}')
    return buf.String()
}

