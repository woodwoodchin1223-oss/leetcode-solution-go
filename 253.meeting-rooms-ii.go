type maxInt []int

func(h *maxInt) Pop() any {
    ret := (*h)[len(*h) - 1]
    (*h) = (*h)[:len(*h) - 1]
    return ret
}

func(h maxInt) Swap(a, b int) {
    h[a], h[b] = h[b], h[a]
}

func(h *maxInt) Push(a any) {
    (*h) = append(*h, a.(int))
}


func(h maxInt) Less(a, b int) bool {
    return h[a] < h[b]
}

func(h maxInt) Len() int {
    return len(h)
}

func minMeetingRooms(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    sort.Slice(intervals, func(a, b int) bool {
        return intervals[a][0] < intervals[b][0]
    })
    h := &maxInt{}
    heap.Init(h)
    heap.Push(h, intervals[0][1])
    for i := 1; i < len(intervals); i++ {
        current := intervals[i]
        if current[0] >= (*h)[0] {
            heap.Pop(h)
        }
        heap.Push(h, current[1])
    }
    return h.Len()
}
