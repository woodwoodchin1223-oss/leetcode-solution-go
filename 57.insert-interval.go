func insert(intervals [][]int, newInterval []int) [][]int {
    ans := make([][]int, 0)
    left := 0
    right := len(intervals) - 1
    for left <= right {
        mid := (left + right) / 2
        if intervals[mid][0] <= newInterval[0] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    toAdd := make([][]int, 0)
    toAdd = append(toAdd, newInterval)

    var nis [][]int
    if left == 0 {
        nis = append(toAdd, intervals...)
    } else if left == len(intervals) {
        nis = append(intervals, toAdd...)
    } else {
        leftList := make([][]int, len(intervals[:left]))
        rightList := make([][]int, len(intervals[left:]))
        copy(leftList, intervals[:left])
        copy(rightList, intervals[left:])
        nis = append(leftList, toAdd...)
        nis = append(nis, rightList...)
        fmt.Println(nis)
    }
    n := len(nis)
    for i := 0; i < n; i++ {
        interval := nis[i]
        if len(ans) == 0 {
            ans = append(ans, interval)
        } else {
            if interval[0] > ans[len(ans) - 1][1] {
                ans = append(ans, interval)
            } else {
                toComp := ans[len(ans) - 1]
                ans = ans[:len(ans) - 1]
                toComp[0] = min(toComp[0], interval[0])
                toComp[1] = max(toComp[1], interval[1])
                ans = append(ans, toComp)
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
