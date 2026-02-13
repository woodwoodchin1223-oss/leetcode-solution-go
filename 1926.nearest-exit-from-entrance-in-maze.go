type Point struct {
    x int
    y int
}

func popLeft(q *[]Point) Point {
    p := (*q)[0]
    *q = (*q)[1:]
    return p
}

func nearestExit(maze [][]byte, entrance []int) int {
    m := len(maze)
    n := len(maze[0])
    queue := make([]Point, 0)
    queue = append(queue, Point{entrance[0], entrance[1]})
    directions := make([][]int, 4)
    directions[0] = []int{0, 1}
    directions[1] = []int{1, 0}
    directions[2] = []int{0, -1}
    directions[3] = []int{-1, 0}
    visited := make(map[Point]int)
    visited[Point{entrance[0], entrance[1]}] = 0
    for len(queue) != 0 {
        current := popLeft(&queue)
        currentStep := visited[current]
        if (current.x == 0 || current.y == 0 || current.x == m - 1 || current.y == n - 1) && current != (Point{entrance[0], entrance[1]}) {
            return currentStep
        }
        for _, d := range(directions) {
            sub := Point{current.x + d[0], current.y + d[1]}
            if _, ok := visited[sub]; ok {
                continue
            }
            if sub.x < 0 || sub.y < 0 || sub.x > m - 1 || sub.y > n - 1 {
                continue
            }
            if maze[sub.x][sub.y] == '+' {
                continue
            } 
            queue = append(queue, sub)
            visited[sub] = currentStep + 1
        }
    }
    return -1
}
