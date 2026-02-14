type Pair struct {
    n1 int
    n2 int
}

type IntHeap []int

func(h IntHeap) Len() int {
    return len(h)
}

func(h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func(h IntHeap) Swap(i, j int) {
    tmp := h[i]
    h[i] = h[j]
    h[j] = tmp
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func max(i, j int64) int64 {
    if i > j {
        return i
    }
    return j
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
    pq := &IntHeap{}
    heap.Init(pq)
    
    pairs := make([]Pair, 0)
    for i := 0; i < len(nums2); i++ {
        n1 := nums1[i]
        n2 := nums2[i]
        pairs = append(pairs, Pair{n1, n2})
    }

    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].n2 > pairs[j].n2
    })

    ans := int64(0)
    tmp := int64(0)
    for i := 0; i < k; i++ {
        tmp += int64(pairs[i].n1)
        heap.Push(pq, pairs[i].n1)
    }
    ans = int64(int64(tmp) * int64(pairs[k - 1].n2))
    for i := k; i < len(nums1); i++ {
        toDelete := int64(heap.Pop(pq).(int))
        tmp += int64(pairs[i].n1) - toDelete
        heap.Push(pq, pairs[i].n1)
        ans = max(tmp * int64(pairs[i].n2), ans)
    }
    return ans
}
