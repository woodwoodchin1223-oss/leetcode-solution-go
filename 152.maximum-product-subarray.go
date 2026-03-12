func maxProduct(nums []int) int {
    large := 1
    small := 1
    tmp := math.MinInt
    for i := 0; i < len(nums); i++ {
        current := nums[i]
        tmpSmall := small
        small = min(current, min(large * current, tmpSmall * current))
        large = max(current, max(large * current, tmpSmall * current))
        tmp = max(tmp, max(current, large))
    }
    return tmp
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
