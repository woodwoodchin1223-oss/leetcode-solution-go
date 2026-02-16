func combinationSum3(k int, n int) [][]int {
    nums := make([]int, 0)
    for i := 1; i <= 9; i++ {
        nums = append(nums, i)
    }
    tmp := make([]int, 0)
    return helper(k, nums, n, tmp)
}

func helper(k int, nums []int, target int, tmp []int) [][]int {
    if k <= 0 {
        return make([][]int, 0)
    }
    if target <= 0 {
        return make([][]int, 0)
    }
    // fmt.Println(k)
    // fmt.Println(target)

    ans := make([][]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i] == target && k == 1 {
            tmp = append(tmp, nums[i])
            b := make([]int, 0)
            for _, t := range tmp {
                b = append(b, t)
            }
            ans = append(ans, b)
            tmp = tmp[:len(tmp)-1]
        } else if nums[i] < target && k > 1 {
            tmp = append(tmp, nums[i])
            sub := helper(k - 1, nums[i+1:], target - nums[i], tmp)
            if len(sub) == 0 {
                tmp = tmp[:len(tmp) - 1]
                continue
            }
            for _, s := range sub {
                if len(s) > 0 {
                    ans = append(ans, s)
                }
            }
            tmp = tmp[:len(tmp) - 1]
        }
    }
    return ans
}
