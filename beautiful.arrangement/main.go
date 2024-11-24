/*
	Leetcode problem: https://leetcode.com/problems/beautiful-arrangement/
*/

package main

func f(i, n int, visitedNum map[int]bool, res *int) {
	if i == n {
		*res++
		return
	}

	for num := 1; num <= n; num++ {
		if !visitedNum[num] && (num%(i+1) == 0 || (i+1)%num == 0) {
			visitedNum[num] = true
			f(i+1, n, visitedNum, res)
			visitedNum[num] = false
		}
	}
}

func countArrangement(n int) int {
	visitedNum := make(map[int]bool)
	var result int

	f(0, n, visitedNum, &result)
	return result
}
