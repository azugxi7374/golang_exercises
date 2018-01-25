package main


func max0(vals ...int) int {
    if len(vals) == 0 {
        return 0
    }

    ret := vals[0]
    for _, val := range vals {
        if ret < val {
            ret = val
        }
    }
    return ret
}

func min0(vals ...int) int {
    if len(vals) == 0 {
        return 0
    }

    ret := vals[0]
    for _, val := range vals {
        if ret > val {
            ret = val
        }
    }
    return ret
}

func max1(x int, vals ...int) int {
    ret := x
    for _, val := range vals {
        if ret < val {
            ret = val
        }
    }
    return ret
}

func min1(x int, vals ...int) int {
    ret := x
    for _, val := range vals {
        if ret > val {
            ret = val
        }
    }
    return ret
}

