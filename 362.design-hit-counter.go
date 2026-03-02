type Pair struct {
    key int
    value int
}
type HitCounter struct {
    total int
    hits []Pair
}


func Constructor() HitCounter {
    hitCounter := HitCounter{}
    hitCounter.hits = make([]Pair, 0)
    hitCounter.total = 0
    return hitCounter
}


func (this *HitCounter) Hit(timestamp int)  {
    this.total += 1
    if len(this.hits) == 0 {
        p := Pair{}
        p.key = timestamp
        p.value = 1
        this.hits = append(this.hits, p)
        return
    } else if this.hits[len(this.hits) - 1].key != timestamp {
        p := Pair{}
        p.key = timestamp
        p.value = 1
        this.hits = append(this.hits, p)
    } else {
        prevCount := this.hits[len(this.hits) - 1].value
        this.hits = this.hits[:len(this.hits) - 1]
        p := Pair{}
        p.key = timestamp
        p.value = prevCount + 1
        this.hits = append(this.hits, p)
    }
}


func (this *HitCounter) GetHits(timestamp int) int {
    for len(this.hits) != 0 {
        diff := timestamp - this.hits[0].key
        if diff >= 300 {
            this.total -= this.hits[0].value
            this.hits = this.hits[1:]
        } else {
            break
        }
    }
    return this.total
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
