type Pair struct {
    first string
    second float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    gid := make(map[string]Pair)
    for i, equation := range equations {
        dividend := equation[0]
        divisor := equation[1]
        v := values[i]
        union(gid, dividend, divisor, v)
    }

    ret := make([]float64, len(queries))
    for i, query := range queries {
        if _, ok := gid[query[0]]; !ok {
            ret[i] = -1.0
            continue
        }
        if _, ok := gid[query[1]]; !ok {
            ret[i] = -1.0
            continue
        }
        ddP := find(gid, query[0])
        dvP := find(gid, query[1])
        if ddP.first != dvP.first {
            ret[i] = -1.0
        } else {
            ret[i] = ddP.second / dvP.second
        }
    }
    return ret
}

func find(gid map[string]Pair, nodeId string) Pair {
    if _, ok := gid[nodeId]; !ok {
        p := Pair{nodeId, 1.0}
        gid[nodeId] = p
    }
    p := gid[nodeId]
    if p.first != nodeId {
        pp := find(gid, p.first)
        ppp := Pair{pp.first, pp.second * p.second}
        gid[nodeId] = ppp
    }
    return gid[nodeId]
}

func union(gid map[string]Pair, dividend string, divisor string, v float64) {
    dividendPair := find(gid, dividend)
    divisorPair := find(gid, divisor)

    ddP := dividendPair.first
    ddV := dividendPair.second
    dvP := divisorPair.first
    dvV := divisorPair.second

    if ddP != dvP {
        p := Pair{dvP, v * dvV / ddV}
        gid[ddP] = p
    }
}
