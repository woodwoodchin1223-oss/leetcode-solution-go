type TrieNode struct {
    isEnd bool
    children []*TrieNode
}

type Trie struct {
    n int
    root *TrieNode
}

func(t *Trie) insert(s string) {
    node := t.root
    for i := 0; i < len(s); i++ {
        current := s[i]
        index := current - 'a'
        if node.children[index] != nil {
            node = node.children[index]
        } else {
            node.children[index] = &TrieNode{false, nil}
             node.children[index].children = make([]*TrieNode, 26)
            node = node.children[index]
        }
    }
    node.isEnd = true
}

func(t *Trie) exist(s string) bool {
    node := t.root
    for _, c := range s {
        if node.children[c] == nil {
            return false
        }
        node = node.children[c]
    }
    return node.isEnd == true
}

func(t *Trie) prefix(s string) bool {
    node := t.root
    for _, c := range s {
        if node.children[c] == nil {
            return false
        }
        node = node.children[c]
    }
    return true
}

func(t *Trie) prefixSearch(s string) *TrieNode {
    node := t.root
    for _, c := range s {
        if node.children[int(c - 'a')] == nil {
            return nil
        }
        node = node.children[int(c - 'a')]
    }
    return node
}

func searchWords(node *TrieNode) []string {
    ret := make([]string, 0)
    for i := 0; i < 26; i++ {
        if node.children[i] != nil {
            sub := searchWords(node.children[i])
            for _, s := range sub {
                ret = append(ret, string('a' + i) + s)
            }
            if node.children[i].isEnd {
                ret = append(ret, string('a' + i))
            }
        }
    }
    return ret
}

func suggestedProducts(products []string, searchWord string) [][]string {
    trie := &Trie{26, nil}
    trie.root = &TrieNode{false, nil}
    trie.root.children = make([]*TrieNode, 26)
    for _, s := range products {
        trie.insert(s)
    }
    ret := make([][]string, 0)
    for i := 0; i < len(searchWord); i++ {
        current := searchWord[0:i+1]
        node := trie.prefixSearch(current)
        if node == nil {
            ret = append(ret, make([]string, 0))
            continue
        }
        collected := searchWords(node)
        for i := 0; i < len(collected); i++ {
            collected[i] = current + collected[i]
        }
        if node.isEnd == true {
            collected = append(collected, current)
        }
        sort.Slice(collected, func(i, j int) bool {
            return collected[i] < collected[j]
        })
        if len(collected) < 3 {
            ret = append(ret, collected)
        } else {
            ret = append(ret, collected[:3])
        }
    }
    return ret
}
