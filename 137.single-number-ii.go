func singleNumber(nums []int) int {
    records := make([]int, 32)
    for _, n := range nums {
        transferToBinary(n, records)
    }
    for i := 0; i < 32; i++ {
        records[i] = records[i] % 3
    }
    fmt.Println(records)
    var res int32 = 0
	for _, b := range records {
		res = (res << 1) | int32(b)
	}
	return int(res)
}

func transferToBinary(n int, records []int) {
    if n == 0 {
		return
	}
	for i := 31; i >= 0; i-- {
		bit := (n >> i) & 1
        records[31 - i] += bit
	}
}
