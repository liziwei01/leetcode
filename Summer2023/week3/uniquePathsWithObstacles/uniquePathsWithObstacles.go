/*
 * @Author: liziwei01
 * @Date: 2023-07-08 15:46:56
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 18:49:02
 * @Description: file content
 */
package uniquepathswithobstacles

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	return recursive(obstacleGrid, m, n, 0, 0)
}

func recursive(obstacleGrid [][]int, m, n, x, y int) int {
	// 超出范围
	if x >= m || y >= n {
		return 0
	}
	// 本格是否为障碍
	if obstacleGrid[x][y] == 1 {
		return 0
	}
	// 走到终点
	if x == m-1 {
		for y < n {
			if obstacleGrid[m-1][y] == 1 {
				return 0
			}
			y++
		}
		return 1
	}
	if y == n-1 {
		for x < m {
			if obstacleGrid[x][n-1] == 1 {
				return 0
			}
			x++
		}
		return 1
	}

	paths := 0
	nextGrid := 0

	// 向右走
	if obstacleGrid[x][y+1] != 1 {
		nextGrid++
		paths += recursive(obstacleGrid, m, n, x, y+2)
	}
	// 向下走
	if obstacleGrid[x+1][y] != 1 {
		nextGrid++
		paths += recursive(obstacleGrid, m, n, x+2, y)
	}
	if nextGrid != 0 {
		paths += nextGrid * recursive(obstacleGrid, m, n, x+1, y+1)
	}

	return paths
}

// import "math/big"

// func uniquePathsWithObstacles(obstacleGrid [][]int) int {
// 	m := len(obstacleGrid)
// 	n := len(obstacleGrid[0])

// 	obstacles := make([][]int, 0)

// 	paths := int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
// 	for i := 0; i != m; i++ {
// 		for j := 0; j != n; j++ {
// 			if obstacleGrid[i][j] == 1 {
// 				paths = paths - comb(i, j)*comb(m+n-2-i-j, n-1-j) + coveredObs(obstacles, i, j)
// 				obstacles = append(obstacles, []int{i, j})
// 			}
// 		}
// 	}

// 	if paths < 0 {
// 		return 0
// 	}

// 	return paths
// }

// func comb(m, n int) int {
// 	return int(new(big.Int).Binomial(int64(m+n), int64(n)).Int64())
// }

// func coveredObs(obstacles [][]int, x, y int) int {
// 	if len(obstacles) == 0 {
// 		return 0
// 	}
// 	paths := 0
// 	for i := 0; i != len(obstacles); i++ {
// 		if x > obstacles[i][0] && y > obstacles[i][1] {
// 			paths += comb()
// 		}
// 	}
// 	return paths
// }
