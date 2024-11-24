/*
	Leetcode problem: https://leetcode.com/problems/word-search/
*/
package main

func outside(i, j, m, n int) bool {
	if (0 <= i && i < m) && (0 <= j && j < n) {
		return false
	}
	return true
}

func dfs(r, c, rows, cols int, word string, visited map[pair]bool, board [][]byte) bool {
	if len(word) == 0 {
		return true
	}

	if outside(r, c, rows, cols) || visited[pair{r: r, c: c}] || board[r][c] != word[0] {
		return false
	}

	visited[pair{r: r, c: c}] = true

	for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		newR, newC := r+d[0], c+d[1]
		if dfs(newR, newC, rows, cols, word[1:], visited, board) {
			return true
		}
	}
	visited[pair{r: r, c: c}] = false
	return false
}

type pair struct {
	r int
	c int
}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	visited := make(map[pair]bool)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, rows, cols, word, visited, board) {
				return true
			}
		}
	}
	return false
}
