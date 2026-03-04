type Node struct {
    key int
    val int
    prev *Node
    next *Node
}

type LRUCache struct {
    capacity int
    count int
    head *Node
    tail *Node
    m map[int]*Node
}

func Constructor(capacity int) LRUCache {
    lruCache := LRUCache{}
    lruCache.head = &Node{}
    lruCache.tail = &Node{}
    lruCache.capacity = capacity
    lruCache.count = 0
    // assign head and tail
    lruCache.head.next = lruCache.tail
    lruCache.tail.prev = lruCache.head
    lruCache.m = make(map[int]*Node)
    return lruCache
}


func (this *LRUCache) Get(key int) int {
    if _, ok := this.m[key]; ok == false {
        return -1
    }
    node, _ := this.m[key]
    nodePrev := node.prev
    nodeNext := node.next

    nodePrev.next = nodeNext
    nodeNext.prev = nodePrev

    this.tail.prev.next = node
    node.prev = this.tail.prev
    node.next = this.tail
    this.tail.prev = node

    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.m[key]; ok == true {
        node, _ := this.m[key]
        node.val = value
        nodePrev := node.prev
        nodeNext := node.next

        nodePrev.next = nodeNext
        nodeNext.prev = nodePrev

        this.tail.prev.next = node
        node.prev = this.tail.prev
        node.next = this.tail
        this.tail.prev = node
    } else {
        if this.count < this.capacity {
            tmp := this.tail.prev
            nodeToAdd := &Node{}
            nodeToAdd.key = key
            nodeToAdd.val = value
            this.m[key] = nodeToAdd
            tmp.next = nodeToAdd
            nodeToAdd.prev = tmp
            nodeToAdd.next = this.tail
            this.tail.prev = nodeToAdd

            this.count += 1
        } else {
            // delete the head
            toRemove := this.head.next
            delete(this.m, toRemove.key)

            tmp := this.head.next.next
            tmp.prev = this.head
            this.head.next = tmp

            
            tmp = this.tail.prev
            nodeToAdd := &Node{}
            nodeToAdd.key = key
            nodeToAdd.val = value
            this.m[key] = nodeToAdd
            tmp.next = nodeToAdd
            nodeToAdd.prev = tmp
            nodeToAdd.next = this.tail
            this.tail.prev = nodeToAdd
        }
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
