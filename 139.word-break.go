func wordBreak(s string, wordDict []string) bool {
    words := make(map[string]bool)
    for _, s := range wordDict {
        words[s] = true
    }
    visited := make(map[string]bool)
    return wordBreakHelper(s, &words, &visited)
}

func wordBreakHelper(s string, words *map[string]bool, visited *map[string]bool) bool {
    if _, ok := (*visited)[s]; ok == true {
        return (*visited)[s]
    }
    if s == "" {
        return true
    }
    for i := 0; i < len(s); i++ {
        current := s[: i + 1]
        if _, ok := (*words)[current]; ok == true {
            sub := wordBreakHelper(s[i + 1:], words, visited)
            if sub == true {
                (*visited)[s] = true
                return true
            }
        }
    }
    (*visited)[s] = false
    return false
}
