/*
 * @Author: liziwei01
 * @Date: 2023-07-13 22:28:33
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-13 22:49:56
 * @Description: file content
 */
package subsets

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(path []int, depth, length, start int)
	dfs = func(path []int, depth, length, start int) {
		if depth == length {
			newPath := make([]int, length)
			copy(newPath, path)
			res = append(res, newPath)
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(path, depth+1, length, i+1)
			path = path[:depth]
		}
	}
	for i := 0; i <= len(nums); i++ {
		dfs(path, 0, i, 0)
	}
	return res
}
