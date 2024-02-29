package main

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	boardmap := make(map[byte]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			boardmap[board[i][j]]++
		}
	}
	for _, c := range word {
		boardmap[byte(c)]--
		if boardmap[byte(c)] < 0 {
			return false
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			visited := make([][]bool, len(board))
			for i := 0; i < len(board); i++ {
				visited[i] = make([]bool, len(board[i]))
			}
			isExist := dfs(board, visited, i, j, []byte{}, word)
			if isExist {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visited [][]bool, i, j int, comp []byte, word string) bool {
	comp = append(comp, board[i][j])
	visited[i][j] = true
	act := string(comp)
	if act == word {
		return true
	}
	if act != word[:len(act)] {
		visited[i][j] = false
		return false
	}
	if i < len(board)-1 && !visited[i+1][j] {
		ans := dfs(board, visited, i+1, j, comp, word)
		if ans {
			return true
		}
	}

	if j < len(board[i])-1 && !visited[i][j+1] {
		ans := dfs(board, visited, i, j+1, comp, word)
		if ans {
			return true
		}
	}

	if i > 0 && !visited[i-1][j] {
		ans := dfs(board, visited, i-1, j, comp, word)
		if ans {
			return true
		}
	}

	if j > 0 && !visited[i][j-1] {
		ans := dfs(board, visited, i, j-1, comp, word)
		if ans {
			return true
		}
	}
	visited[i][j] = false
	return false
}
