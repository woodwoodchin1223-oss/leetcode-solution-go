type Node struct {
    key int
    val int
    prev *Node
    next *Node
}

type LRUCache struct {
    capacity int
    head *Node
    tail *Node
    m map[int]*Node
}

func Constructor(capacity int) LRUCache {
    lruCache := LRUCache{}
    lruCache.head = &Node{}
    lruCache.tail = &Node{}
    lruCache.capacity = capacity
    // assign head and tail
    lruCache.head.next = lruCache.tail
    lruCache.tail.prev = lruCache.head
    lruCache.m = make(map[int]*Node, 0)
    return lruCache
}


func (this *LRUCache) Get(key int) int {
    
}


func (this *LRUCache) Put(key int, value int)  {
    
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
