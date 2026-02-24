func minEatingSpeed(piles []int, h int) int {
    left := 1
    right := 0
    for _, p := range piles {
        right = max(right, p)
    }
    for left <= right {
        mid := (left + right) / 2
        if calc(mid, piles) > h {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func calc(speed int, piles []int) int {
    ans := 0
    for _, p := range piles {
        if (p / speed) * speed < p {
            ans += p / speed + 1
        } else {
            ans += p / speed
        }
    }
    return ans
}
