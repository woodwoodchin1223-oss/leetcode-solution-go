func firstMissingPositive(nums []int) int {
    n := len(nums)
    contains1 := false
    
    for i := 0; i < n; i++ {
        if nums[i] == 1 {
            contains1 = true
        }
        if nums[i] <= 0 || nums[i] > n {
            nums[i] = 1
        }
    }

    if contains1 != true {
        return 1
    }

    for i := 0; i < n; i++ {
        value := abs(nums[i])
        if value == n {
            nums[0] = -1 * abs(nums[0])
        } else {
            nums[value] = -1 * abs(nums[value])
        }
    }

    for i := 1; i < n; i++ {
        if nums[i] > 0 {
            return i
        }
    }
    
    if nums[0] > 0 {
        return n
    }

    return n + 1
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -1 * a
}
