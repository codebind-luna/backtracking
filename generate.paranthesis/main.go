/*
	Leetcode problem: https://leetcode.com/problems/generate-parentheses/
*/
package main

import "strings"

func allPermutations(n, open, close int, ds []string, ans *[]string) {
	if len(ds) == 2*n {
		*ans = append(*ans, strings.Join(ds, ""))
		return
	}

	if open < n {
		ds = append(ds, "(")
		allPermutations(n, open+1, close, ds, ans)
		ds = ds[:len(ds)-1]
	}

	if open > close {
		ds = append(ds, ")")
		allPermutations(n, open, close+1, ds, ans)
		ds = ds[:len(ds)-1]
	}
}

func generateParenthesis(n int) []string {
	var ans []string
	var ds []string

	allPermutations(n, 0, 0, ds, &ans)
	return ans
}
