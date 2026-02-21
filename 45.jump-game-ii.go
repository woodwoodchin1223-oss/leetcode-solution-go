func jump(nums []int) int {
    currEnd := 0
    currFar := 0
    ans := 0
    for i := 0; i < len(nums) - 1; i++ {
        currFar = max(currFar, i + nums[i])
        if i == currEnd {
            currEnd = currFar
            ans += 1
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
