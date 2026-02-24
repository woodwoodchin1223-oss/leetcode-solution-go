func minCost(costs [][]int) int {
    dp := make([][]int, 0)
    for i := 0; i < 3; i++ {
        dp = append(dp, make([]int, len(costs)))
    }
    if len(costs) == 1 {
        return min(costs[0][0], min(costs[0][1], costs[0][2]))
    }
    sub1 := costs[0][0] + minCostHelper(costs, dp, 0, 1)
    sub2 := costs[0][1] + minCostHelper(costs, dp, 1, 1)
    sub3 := costs[0][2] + minCostHelper(costs, dp, 2, 1)
    return min(sub1, min(sub2, sub3))
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minCostHelper(costs [][]int, dp [][]int, prev int, index int) int {
    if dp[prev][index] != 0 {
        return dp[prev][index]
    }
    if prev == 1 {
        if index == len(costs) - 1 {
            return min(costs[index][0], costs[index][2])
        } else {
            sub1 := costs[index][2] + minCostHelper(costs, dp, 2, index + 1)
            sub2 := costs[index][0] + minCostHelper(costs, dp, 0, index + 1)
            dp[prev][index] = min(sub1, sub2)
            return dp[prev][index]
        }
    } else if prev == 0 {
        if index == len(costs) - 1 {
            return min(costs[index][1], costs[index][2])
        } else {
            sub1 := costs[index][2] + minCostHelper(costs, dp, 2, index + 1)
            sub2 := costs[index][1] + minCostHelper(costs, dp, 1, index + 1)
            dp[prev][index] = min(sub1, sub2)
            return dp[prev][index]
        }
    } else {
        if index == len(costs) - 1 {
            return min(costs[index][1], costs[index][0])
        } else {
            sub1 := costs[index][0] + minCostHelper(costs, dp, 0, index + 1)
            sub2 := costs[index][1] + minCostHelper(costs, dp, 1, index + 1)
            dp[prev][index] = min(sub1, sub2)
            return dp[prev][index]
        }
    }
}
