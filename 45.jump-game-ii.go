func jump(nums []int) int {
    answer := 0
    n := len(nums)
    curEnd := 0
    curFar := 0

    for i := 0; i < n - 1; i++ {
        curFar = max(curFar, nums[i] + i)
        if i == curEnd {
            answer += 1
            curEnd = curFar
        }
    }
    return answer
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
