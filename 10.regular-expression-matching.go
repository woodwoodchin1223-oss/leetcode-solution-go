func isMatch(s string, p string) bool {
    dp := make([][]bool, 0)
    for i:= 0; i < len(s) + 1; i++ {
        dp = append(dp, make([]bool, len(p) + 1))
    }
    for i := 0; i < len(s) + 1; i++ {
        for j := 0; j < len(p) + 1; j++ {
            dp[i][j] = false
        }
    }
    dp[0][0] = true

    for i := 1; i < len(p) + 1; i++ {
        if (p[i - 1] == '*') {
            dp[0][i] = dp[0][i - 2]
        } 
    }

    for i := 1; i < len(s) + 1; i++ {
        for j := 1; j < len(p) + 1; j++ {
            if (s[i - 1] == p[j - 1] || p[j - 1] == '.') {
                dp[i][j] = dp[i - 1][j - 1]
            } else {
                if (p[j - 1] == '*' && (p[j - 2] == '.' || p[j - 2] == s[i - 1])) {
                    dp[i][j] = dp[i][j - 1] || dp[i][j - 2] || dp[i - 1][j]
                } else if p[j - 1] == '*' && (p[j-2] != s[i-1] && p[j-2] != '.') {
                    dp[i][j] = dp[i][j - 2]
                }
            }
        }
    }
    return dp[len(s)][len(p)]
}
