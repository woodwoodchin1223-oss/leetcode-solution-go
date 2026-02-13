type Point struct {
    x int
    y int
}

func popLeft(queue *[]Point) Point {
    q := (*queue)[0]
    (*queue) = (*queue)[1:]
    return q
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func orangesRotting(grid [][]int) int {
    visited := make(map[Point]int)
    queue := make([]Point, 0)
    total := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 2 {
                queue = append(queue, Point{i, j})
                visited[Point{i, j}] = 0
            }
            if grid[i][j] != 0 {
                total += 1
            }
        }
    }
    ans := 0
    directions := make([]Point, 4)
    directions = append(directions, Point{1, 0})
    directions = append(directions, Point{0, 1})
    directions = append(directions, Point{0, -1})
    directions = append(directions, Point{-1, 0})

    calc := 0
    for len(queue) != 0 {
        current := popLeft(&queue)
        calc += 1
        time := visited[current]
        ans = max(ans, time)
        for _, d := range directions {
            sub := Point{current.x + d.x, current.y + d.y}
            if _, ok := visited[sub]; ok {
                continue
            }
            if sub.x < 0 || sub.y < 0 || sub.x >= len(grid) || sub.y >= len(grid[0]) {
                continue
            }
            if grid[sub.x][sub.y] == 0 || grid[sub.x][sub.y] == 2 {
                continue
            }
            queue = append(queue, sub)
            visited[sub] = time + 1
        }
    }

    if calc != total {
        return -1
    }
    return ans
}
