func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func maxArea(height []int) int {
    ans := 0

    left := 0
    right := len(height) - 1

    for left < right {
        ans = max(ans, min(height[left], height[right]) * (right - left))
        if height[left] < height[right] {
            left += 1
        } else {
            right -= 1
        }
    }
    return ans
}
