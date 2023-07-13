/*
 * @Author: liziwei01
 * @Date: 2023-07-11 01:55:36
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-11 20:41:58
 * @Description: file content
 */
package searchmatrix

// 二分查找
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	slice := make([]int, 0)
	for i := 0; i != m; i++ {
		slice = append(slice, matrix[i]...)
	}

	return binarySearch(slice, target)
}
func binarySearch(slice []int, target int) bool {
	n := len(slice)
	if n == 0 {
		return false
	}
	if slice[n/2] == target {
		return true
	}
	if slice[n/2] > target {
		return binarySearch(slice[:n/2], target)
	}
	return binarySearch(slice[n/2+1:], target)
}
