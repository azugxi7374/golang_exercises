package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount0(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountEx3(x uint64) int {
    var ret byte = 0
    for i:=0; i<8; i++ {
        ret += pc[byte(x>>uint(i*8))]
    }
    return int(ret)
}

func PopCountEx4(x uint64) int {
    var ret = byte(0)
    for i:=uint(0); i<64; i++ {
        ret += byte(x&1)
        x>>=1
    }
    return int(ret)
}

func PopCountEx5(x uint64) int {
    cnt := 0
    for x != 0 {
        x &= x-1
        cnt++
    }
    return cnt
}
