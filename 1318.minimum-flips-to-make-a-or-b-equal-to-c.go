func reverseString(s string) string {
    bytes := []byte(s)
    n := len(bytes)
    for i := 0; i < n/2; i++ {
        bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
    }
    return string(bytes)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -1 * a
}

func minFlips(a int, b int, c int) int {
    a1 := reverseString(strconv.FormatInt(int64(a), 2))
    b1 := reverseString(strconv.FormatInt(int64(b), 2))
    c1 := reverseString(strconv.FormatInt(int64(c), 2))
    maxn := max(len(a1), max(len(b1), len(c1)))
    for i := 0; i < maxn; i++ {
        if i >= len(a1) {
            a1 += "0"
        }
        if i >= len(b1) {
            b1 += "0"
        }
        if i >= len(c1) {
            c1 += "0"
        }
    }
    ret := 0
    for i := 0; i < maxn; i++ {
        a11 := int(a1[i] - '0')
        b11 := int(b1[i] - '0')
        c11 := int(c1[i] - '0')
        if (a11 == 1 && c11 == 1) || (b11 == 1 && c11 == 1) {
            ret += 0
        } else {
            if a11 == 0 && b11 == 0 && c11 == 1 {
                ret += 1
            } else if a11 == 0 && b11 == 0 && c11 == 0 {
                ret += 0
            } else {
                ret += abs(a11 + b11 - c11)
            }
        }
    }
    return ret
}
