func twoSum(nums []int, target int) []int {
    ans := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if tmp := nums[i] + nums[j]; tmp == target {
                ans = append(ans, i)
                ans = append(ans, j)
                return ans
            }
        }
    }
    return ans
}
