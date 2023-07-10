/*
 * @Author: liziwei01
 * @Date: 2023-07-08 15:18:52
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 16:45:28
 * @Description: file content
 */
package uniquepaths

// var (
// 	dict = make(map[int]int, 0)
// )

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	return 2*uniquePaths(m-1, n-1) + uniquePaths(m-2, n) + uniquePaths(m, n-2)
}

// func hasPath(obstacleGrid [][]int, m, n, i, j int) bool {
// 	if i >= m || j >= n {
// 		return true
// 	}
// 	if obstacleGrid[i][j] == 1 {
// 		return false
// 	}

// 	return hasPath(obstacleGrid, m, n, i+1, j) || hasPath(obstacleGrid, m, n, i, j+1)
// }

// func recursive(m int, n int) int {
// 	dict[1] = 1
// 	dict[2] = 2
// 	dict[3] = 6
// 	dict[4] = 20
// 	dict[5] = 70
// 	dict[6] = 252
// 	dict[7] = 924
// 	dict[8] = 3432
// 	dict[9] = 12870
// 	if m == n && m <= 9 {
// 		return dict[m]
// 	}
// 	if m == 1 || n == 1 {
// 		return 1
// 	}

// 	return recursive(m-1, n) + recursive(m, n-1)
// }
