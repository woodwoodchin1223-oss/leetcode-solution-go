func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    sort.Slice(nums, func(a, b int) bool {
        return nums[a] < nums[b]
    })

    longestSteak := 1
    currentStreak := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i - 1] {
            if nums[i] == nums[i - 1] + 1 {
                currentStreak += 1
            } else {
                if currentStreak > longestSteak {
                    longestSteak = currentStreak
                }
                currentStreak = 1
            }
        } 
    }
    if currentStreak > longestSteak {
        longestSteak = currentStreak
    }
    return longestSteak
}
