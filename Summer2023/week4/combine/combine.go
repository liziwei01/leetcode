/*
 * @Author: liziwei01
 * @Date: 2023-07-13 21:54:25
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-13 22:26:16
 * @Description: file content
 */
package combine

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(path []int, depth, start int)
	dfs = func(path []int, depth, start int) {
		if depth == k {
			newPath := make([]int, k)
			copy(newPath, path)
			res = append(res, newPath)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(path, depth+1, i+1)
			path = path[:depth]
		}
	}
	dfs(path, 0, 1)
	return res
}