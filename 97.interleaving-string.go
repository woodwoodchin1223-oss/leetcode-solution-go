func isInterleave(s1 string, s2 string, s3 string) bool {
    dp := make([][]int, 0)
    if len(s1) + len(s2) != len(s3) {
        return false
    }
    for i := 0; i < len(s1); i++ {
        dp = append(dp, make([]int, 0))
        for j := 0; j < len(s2); j++ {
            dp[i] = append(dp[i], -1)
        }
    }
    ret := helper(dp, s1, s2, s3, 0, 0, 0)
    if ret == 1 {
        return true
    }
    return false
}

func helper(dp [][]int, s1 string, s2 string, s3 string, i int, j int, k int) int {
    if i == len(s1) {
        if s2[j:] == s3[k:] {
            return 1
        } else {
            return 0
        }
    }
    if j == len(s2) {
        if s1[i:] == s3[k:] {
            return 1
        } else {
            return 0
        }
    }
    
    if dp[i][j] != -1 {
        return dp[i][j]
    }
    sub1 := 0
    if s1[i] == s3[k] {
        sub1 = helper(dp, s1, s2, s3, i + 1, j, k + 1)
    }
    sub2 := 0
    if s2[j] == s3[k] {
         sub2 = helper(dp, s1, s2, s3, i, j + 1, k + 1)
    }
    dp[i][j] = 0
    if sub1 == 1 || sub2 == 1 {
        dp[i][j] = 1
    }
    return dp[i][j]
}
