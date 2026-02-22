func canJump(nums []int) bool {
    curEnd := 0
    curFar := 0
    if len(nums) == 1 {
        return true
    }
    for i:= 0; i < len(nums) - 1; i++ {
        curFar = max(curFar, i + nums[i])
        if curEnd == i {
            curEnd = curFar
            if curEnd >= len(nums) - 1 {
                return true
            }
        }
    }
    return false
}

func max(a ,b int) int {
    if a > b {
        return a
    }
    return b
}
