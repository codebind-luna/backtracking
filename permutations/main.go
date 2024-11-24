/*
	Leetcode problem: https://leetcode.com/problems/permutations/
*/
package main

func f(i int, nums, ds []int, visited map[int]bool, o *[][]int) {
	if i == len(nums) {
		*o = append(*o, append([]int{}, ds...))
		return
	}

	for j := 0; j < len(nums); j++ {
		if !visited[nums[j]] {
			visited[nums[j]] = true
			ds = append(ds, nums[j])
			f(i+1, nums, ds, visited, o)
			ds = ds[:len(ds)-1]
			visited[nums[j]] = false
		}
	}
}

func permute(nums []int) [][]int {
	visited := make(map[int]bool)
	var result [][]int
	f(0, nums, []int{}, visited, &result)
	return result
}
