var ans int
var sans string
func longestPalindrome(s string) string {
    n := len(s)
    ans = 1
    sans = string(s[0])
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            dp[i][j] = -1
        }
    }
    for i := 0; i < n; i++ {
        dp[i][i] = 1
    }
    helper(dp, 0, len(s) - 1, s)
    return sans
}

func helper(dp [][]int, i int, j int, s string) int {
    if dp[i][j] != -1 {
        return dp[i][j]
    }
    if i + 1 == j {
        if s[i] == s[j] {
            dp[i][j] = 2
            if dp[i][j] > ans {
                ans = dp[i][j]
                sans = s[i : j+1]
            }
            return 2
        } else {
            dp[i][j] = 0
            return 0
        }
    } else {
        if s[i] == s[j] {
            sub := helper(dp, i + 1, j - 1, s)
            if sub != j - i - 1 {
                helper(dp, i, j - 1, s)
                helper(dp, i + 1, j, s)
                dp[i][j] = 0
                return 0
            } else {
                dp[i][j] = j - i + 1
                if dp[i][j] > ans {
                    ans = dp[i][j]
                    sans = s[i : j+1]
                }
                return dp[i][j]
            }
        } else {
            helper(dp, i + 1, j - 1, s)
            helper(dp, i, j - 1, s)
            helper(dp, i + 1, j, s)
            dp[i][j] = 0
            return 0
        }
    }
}
