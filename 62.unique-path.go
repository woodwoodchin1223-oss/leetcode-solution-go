
func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dp[i][j] = -1
        }
    }
    return helper(dp, 0, 0, m, n)
}

func helper(dp [][]int, row int, col int, m int, n int) int {
    if row >= m || col >= n {
        return -1
    }

    if dp[row][col] != -1 {
        return dp[row][col]
    }

    if row == m - 1 && col == n - 1 {
        return 1
    }

    h1 := helper(dp, row + 1, col, m, n)
    h2 := helper(dp, row, col + 1, m, n)
    if h1 == -1 {
        dp[row][col] = h2
    } else if h2 == -1 {
        dp[row][col] = h1
    } else {
        dp[row][col] = h1 + h2
    }
    return dp[row][col]
}
