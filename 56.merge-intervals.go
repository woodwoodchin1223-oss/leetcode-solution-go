func merge(intervals [][]int) [][]int {
    ans := make([][]int, 0)
    sort.Slice(intervals, func(a, b int) bool {
        return intervals[a][0] <= intervals[b][0]
    })
    for i := 0; i < len(intervals); i++ {
        interval := intervals[i]
        if len(ans) == 0 {
            ans = append(ans, interval)
        } else {
            if ans[len(ans) - 1][1] < interval[0] {
                ans = append(ans, interval)
            } else {
                toComp := ans[len(ans) - 1]
                ans = ans[:len(ans) - 1]
                ans = append(ans, []int{min(toComp[0], interval[0]), max(toComp[1], interval[1])})
            }
        }
    }
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
