/*
 * @Author: liziwei01
 * @Date: 2023-06-28 10:49:56
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-04 17:43:11
 * @Description: file content
 */
package rotate

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}	
}