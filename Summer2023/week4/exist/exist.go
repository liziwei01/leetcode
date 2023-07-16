/*
 * @Author: liziwei01
 * @Date: 2023-07-15 02:18:51
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-15 02:47:21
 * @Description: file content
 */
package exist

type int2 struct {
	a int
	b int
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make(map[int2]bool)
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i-1 >= 0 && !used[int2{i-1, j}] && board[i-1][j] == word[k] {
			used[int2{i-1, j}] = true
			if dfs(i-1, j, k+1) {
				return true
			} else {
				used[int2{i-1, j}] = false
			}
		}
		if j-1 >= 0 && !used[int2{i, j-1}] && board[i][j-1] == word[k] {
			used[int2{i, j-1}] = true
			if dfs(i, j-1, k+1) {
				return true
			} else {
				used[int2{i, j-1}] = false
			}
		}
		if i+1 < m && !used[int2{i+1, j}] && board[i+1][j] == word[k] {
			used[int2{i+1, j}] = true
			if dfs(i+1, j, k+1) {
				return true
			} else {
				used[int2{i+1, j}] = false
			}
		}
		if j+1 < n && !used[int2{i, j+1}] && board[i][j+1] == word[k] {
			used[int2{i, j+1}] = true
			if dfs(i, j+1, k+1) {
				return true
			} else {
				used[int2{i, j+1}] = false
			}
		}

		return false
	}

	for i := 0; i != m; i++ {
		for j := 0; j != n; j++ {
			if board[i][j] == word[0] {
				used[int2{i, j}] = true
				if dfs(i, j, 1) {
					return true
				} else {
					used[int2{i, j}] = false
				}
			}
		}
	}

	return false
}
