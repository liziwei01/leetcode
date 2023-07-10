/*
 * @Author: liziwei01
 * @Date: 2023-07-08 23:20:58
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 23:33:28
 * @Description: file content
 */
package minpathsum

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	sum := make([][]int, 0)
	for i := 0; i != m; i++ {
		temp := make([]int, n)
		sum = append(sum, temp)
	}
	
	for i := 0; i != m; i++ {
		for j := 0; j != n; j++ {
			if i != 0 && j != 0 {
				sum[i][j] = min(sum[i-1][j]+grid[i][j], sum[i][j-1]+grid[i][j])
			} else if i != 0 && j == 0 {
				sum[i][j] = sum[i-1][j]+grid[i][j]
			} else if i == 0 && j != 0 {
				sum[i][j] = sum[i][j-1]+grid[i][j]
			} else {
				sum[i][j] = grid[i][j]
			}
		}
	}

	return sum[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}