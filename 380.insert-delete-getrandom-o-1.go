type RandomizedSet struct {
    index map[int]int
    container []int
}


func Constructor() RandomizedSet {
    randomizedSet := RandomizedSet{}
    randomizedSet.index = make(map[int]int)
    randomizedSet.container = make([]int, 0)
    return randomizedSet
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.index[val]; ok {
        return false
    }
    current := len(this.container)
    this.container = append(this.container, val)
    this.index[val] = current
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.index[val]; ok == false {
        return false
    }

    current, _ := this.index[val]
    if len(this.container) == 1 {
        this.container = make([]int, 0)
        this.index = make(map[int]int)
        return true
    }
    if current == len(this.container) - 1 {
        this.container = this.container[:len(this.container) - 1]
        delete(this.index, val)
        return true
    }
    toInsert := this.container[len(this.container) - 1]
    this.container[current] = toInsert
    this.index[toInsert] = current
    delete(this.index, val)
    this.container = this.container[:len(this.container) - 1]
    return true
}

func (this *RandomizedSet) GetRandom() int {
    x := rand.Intn(len(this.container))
    return this.container[x]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
