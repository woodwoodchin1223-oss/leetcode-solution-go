func longestOnes(nums []int, k int) int {
    left := 0
    right := 0
    ans := 0
    for left <= right && right < len(nums) {
        if nums[right] == 1 {
            right += 1
            ans = max(ans, right - left)
        } else {
            k -= 1
            right += 1
            for k < 0 {
                if nums[left] == 0 {
                    k += 1
                } 
                left += 1
            }
            ans = max(ans, right - left)
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
