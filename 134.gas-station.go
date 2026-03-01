func canCompleteCircuit(gas []int, cost []int) int {
    if len(gas) == 1 {
        if gas[0] >= cost[0] {
            return 0
        }
        return -1
    }
    oil := gas[0]
    start := 0
    i := 1
    sign := 0
    n := len(gas)
    for true {
        index := -1
        if i - 1 < 0 {
            index = n - 1
        } else {
            index = i - 1
        }
        oil -= cost[index]
        if oil < 0 {
            start = i
            if sign == 1 {
                return -1
            }
            oil = gas[i]
        } else {
            oil += gas[i]
        }
        i = (i + 1) % n
        if i == 0 {
            sign += 1
        }
        if sign == 2 {
            fmt.Println(start)
            break
        }
    }
    return start
}
