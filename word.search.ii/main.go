/*
	Leetcode problem: https://leetcode.com/problems/word-search-ii/
*/
package main

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func Construct() *Trie {
	return &Trie{}
}

func (this *Trie) add(s string) {
	curr := this

	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = Construct()
		}
		curr = curr.children[idx]
	}

	curr.isEnd = true
}
func outside(i, j, m, n int) bool {
	if (0 <= i && i < m) && (0 <= j && j < n) {
		return false
	}
	return true
}

func dfs(r, c, rows, cols int, node *Trie, word string, visited map[pair]bool, board [][]byte, result map[string]bool) {
	if outside(r, c, rows, cols) || visited[pair{r: r, c: c}] || node.children[board[r][c]-'a'] == nil {
		return
	}

	visited[pair{r: r, c: c}] = true
	node = node.children[board[r][c]-'a']
	word += string(board[r][c])
	if node.isEnd {
		result[word] = true
	}
	for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		newR, newC := r+d[0], c+d[1]
		dfs(newR, newC, rows, cols, node, word, visited, board, result)
	}
	visited[pair{r: r, c: c}] = false
}

type pair struct {
	r int
	c int
}

func findWords(board [][]byte, words []string) []string {
	root := Construct()

	for _, word := range words {
		root.add(word)
	}
	rows := len(board)
	cols := len(board[0])
	ans := make(map[string]bool)

	visited := make(map[pair]bool)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, rows, cols, root, "", visited, board, ans)
		}
	}

	res := []string{}
	for k := range ans {
		res = append(res, k)
	}
	return res
}
