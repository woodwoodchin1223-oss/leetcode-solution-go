func exist(board [][]byte, word string) bool {
    directions := make([][]int, 0)
    directions = append(directions, []int{0, 1})
    directions = append(directions, []int{1, 0})
    directions = append(directions, []int{0, -1})
    directions = append(directions, []int{-1, 0})
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            sub := traverse(board, i, j, word, 0)
            if sub == true {
                return true
            }
        }
    }
    return false
}

func traverse(board [][]byte, i int, j int, word string, index int) bool {
    if index >= len(word) {
        return true
    }
    

    if i < 0 || i >= len(board) {
        return false
    }

    if j < 0 || j >= len(board[0]) {
        return false
    }

    if board[i][j] == '.' {
        return false
    }

    if board[i][j] != word[index] {
        return false
    }

    currentChar := board[i][j]
    board[i][j] = '.'
    directions := make([][]int, 0)
    directions = append(directions, []int{0, 1})
    directions = append(directions, []int{1, 0})
    directions = append(directions, []int{0, -1})
    directions = append(directions, []int{-1, 0})
    for _, d := range directions {
        sub := traverse(board, i + d[0], j + d[1], word, index + 1)
        if sub == true {
            return true
        }
    }
    board[i][j] = currentChar
    return false
    
}
