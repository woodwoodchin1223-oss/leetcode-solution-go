func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
func lengthOfLongestSubstring(s string) int {
    m := make(map[byte]int)
    left := 0
    right := 0
    ans := 0
    tmp := 0
    for ;left <= right && right < len(s); {
        if _, ok := m[s[right]]; ok {
            for ok {
                m[s[left]] -= 1
                if m[s[left]] == 0 {
                    delete(m, s[left])
                }
                left += 1
                tmp -= 1
                _, ok = m[s[right]]
            }
        } else {
            m[s[right]] = 1
            tmp += 1
            right += 1
            ans = max(ans, tmp)
        }
    }
    return ans
}
