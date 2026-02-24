type MinStack struct {
    minVal int
    minNums []int
    nums []int
}


func Constructor() MinStack {
    minStack := MinStack{}
    minStack.minNums = make([]int, 0)
    minStack.nums = make([]int, 0)
    minStack.minVal = math.MaxInt
    return minStack
}


func (this *MinStack) Push(val int)  {
    if val < this.minVal {
        this.minVal = val
    }
    this.minNums = append(this.minNums, this.minVal)
    this.nums = append(this.nums, val)
}


func (this *MinStack) Pop()  {
    this.nums = this.nums[:len(this.nums) - 1]
    this.minNums = this.minNums[:len(this.minNums) - 1]
    if len(this.nums) == 0 {
        this.minVal = math.MaxInt
    } else {
        if this.minNums[len(this.minNums) - 1] > this.minVal {
            this.minVal = this.minNums[len(this.minNums) - 1]
        }
    }
}


func (this *MinStack) Top() int {
    return this.nums[len(this.nums) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minNums[len(this.minNums) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
