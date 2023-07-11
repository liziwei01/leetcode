/*
 * @Author: liziwei01
 * @Date: 2023-07-11 00:44:40
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-11 01:08:34
 * @Description: file content
 */
package setzeroes

func setZeroes(matrix [][]int) {
	zeroes := make([][]int, 0)
	m, n := len(matrix), len(matrix[0])
	for i := 0; i != m; i++ {
		for j := 0; j != n; j++ {
			if matrix[i][j] == 0 {
				zeroes = append(zeroes, []int{i, j})
			}
		}
	}
	for i := 0; i != len(zeroes); i++ {
		x, y := zeroes[i][0], zeroes[i][1]
		for j := 0; j != m; j++ {
			matrix[j][y] = 0
		}
		for k := 0; k != n; k++ {
			matrix[x][k] = 0
		}
	}
}
