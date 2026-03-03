func majorityElement(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }
    num1 := nums[0]
    num2 := nums[1]
    count1 := 0
    count2 := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == num1 {
            count1 += 1
        } else if nums[i] == num2 {
            count2 += 1
        } else if count1 == 0 {
            count1 += 1
            num1 = nums[i]
        } else if count2 == 0 {
            count2 += 1
            num2 = nums[i]
        } else {
            count1 -= 1
            count2 -= 1
        }
    }
    
    count1 = 0
    count2 = 0
    for _, n := range nums {
        if n == num1 {
            count1 += 1
        }
        if n == num2 {
            count2 += 1
        }
    }
    
    ans := make([]int, 0)
    if count1 > len(nums) / 3 {
        ans = append(ans, num1)
    }
    if count2 > len(nums) / 3 && num2 != num1 {
        ans = append(ans, num2)
    }
    return ans
}
