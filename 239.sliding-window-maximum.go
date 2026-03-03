func maxSlidingWindow(nums []int, k int) []int {
    stack := make([]int, 0)
    for i := 0; i < min(k, len(nums)); i++ {
        if len(stack) == 0 {
            stack = append(stack, i) 
        } else {
            for len(stack) > 0 && nums[stack[len(stack) - 1]] <= nums[i] {
                stack = stack[:len(stack) - 1]
            }
            stack = append(stack, i)
        }
    }
    
    ans := make([]int, 0)
    ans = append(ans, nums[stack[0]])

    for i := k; i < len(nums); i++ {
        for len(stack) > 0 && stack[0] <= i - k {
            stack = stack[1:len(stack)]
        }
        for len(stack) > 0 && nums[stack[len(stack) - 1]] <= nums[i] {
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, i)
        ans = append(ans, nums[stack[0]])
    }
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
