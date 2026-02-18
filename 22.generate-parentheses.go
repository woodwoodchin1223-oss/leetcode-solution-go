func helper(left, right int, tmp string, ans *[]string, visited map[string]bool) {
    if left == 0 && right == 0 {
        if _, ok := visited[tmp];!ok {
            *ans = append(*ans, tmp)
            visited[tmp] = true
        }
        return
    }
    if left > 0 {
        tmp += "("
        helper(left - 1, right, tmp, ans, visited)
        tmp = tmp[:len(tmp) - 1]
    }
    
    if right > 0 && left < right {
        tmp += ")"
        helper(left, right - 1, tmp, ans, visited)
        tmp = tmp[:len(tmp) - 1]
    }
}

func generateParenthesis(n int) []string {
    ans := make([]string, 0)
    helper(n, n, "", &ans, make(map[string]bool))
    return ans
}
