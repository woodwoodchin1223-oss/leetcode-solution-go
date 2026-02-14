type IntHeap []int

func(h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func(h IntHeap) Len() int {
    return len(h)
}

func(h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func(h *IntHeap) Pop() any {
    n := len(*h)
    x := (*h)[n - 1]
    *h = (*h)[:n - 1]
    return x
}

func(h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func totalCost(costs []int, k int, candidates int) int64 {
    left := &IntHeap{}
    right := &IntHeap{}
    heap.Init(left)
    heap.Init(right)
    for i := 0; i < candidates; i++ {
        heap.Push(left, costs[i])
    }

    for i := max(candidates, len(costs)-candidates); i < len(costs); i++ {
        heap.Push(right, costs[i])
    }

    leftIndex := candidates
    rightIndex := len(costs) - candidates - 1
    totalCost := int64(0)

    count := 0
    
    for leftIndex <= rightIndex &&  count < k {
        leftHead := heap.Pop(left).(int)
        rightHead := heap.Pop(right).(int)
        fmt.Println(leftHead)
        fmt.Println(rightHead)
        fmt.Println("!!!!")
        count += 1
        if leftHead <= rightHead {
            totalCost += int64(leftHead)
            heap.Push(right, rightHead)
            heap.Push(left, costs[leftIndex])
            leftIndex += 1
        } else {
            totalCost += int64(rightHead)
            heap.Push(right, costs[rightIndex])
            heap.Push(left, leftHead)
            rightIndex -= 1
        }
    }
    
    if count == k {
        return totalCost
    }

    for count < k {
        count += 1
        if len(*left) == 0 && len(*right) == 0 {
            return totalCost
        } else if len(*left) == 0 && len(*right) > 0 {
            rightHead := heap.Pop(right).(int)
            totalCost += int64(rightHead)
        } else if len(*right) == 0 && len(*left) > 0 {
            leftHead := heap.Pop(left).(int)
            totalCost += int64(leftHead)
        } else {
            leftHead := heap.Pop(left).(int)
            rightHead := heap.Pop(right).(int)

            if leftHead <= rightHead {
                
                totalCost += int64(leftHead)
                heap.Push(right, rightHead)
            } else {
                
                totalCost += int64(rightHead)
                heap.Push(left, leftHead)
            }
        }
    }
    return totalCost

}
