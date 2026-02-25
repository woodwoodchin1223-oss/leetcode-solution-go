func maxSubArray(nums []int) int {
    ans := 0
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    ans = max(dp[0], ans)
    for i := 1; i < len(nums); i++ {
        dp[i] = max(dp[i - 1] + nums[i], nums[i]) 
        ans = max(dp[i], ans)
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
