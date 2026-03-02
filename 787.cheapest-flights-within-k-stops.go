func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    dist := make([]int, 0)
    for i := 0; i < n; i++ {
        dist = append(dist, math.MaxInt)
    }
    dist[src] = 0
    for i := 0; i <= k; i++ {
        temp := make([]int, n)
        copy(temp, dist)
        for _, flight := range flights {
            if dist[flight[0]] != math.MaxInt {
                temp[flight[1]] = min(temp[flight[1]], dist[flight[0]] + flight[2])
            }
        }
        dist = temp
    }
    if dist[dst] == math.MaxInt {
        return -1
    }
    return dist[dst]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
