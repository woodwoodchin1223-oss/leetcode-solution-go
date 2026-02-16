func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]
    })
    ret := 1
    firstEnd := points[0][1]
    for i := 1; i < len(points); i++ {
        point := points[i]
        xStart := point[0]
        xEnd := point[1]
        if firstEnd < xStart {
            ret += 1
            firstEnd = xEnd
        }
    }
    return ret
}
