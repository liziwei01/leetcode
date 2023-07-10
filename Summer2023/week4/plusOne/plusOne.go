/*
 * @Author: liziwei01
 * @Date: 2023-07-09 00:34:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-09 00:46:14
 * @Description: file content
 */
package plusone

func plusOne(digits []int) []int {
	lastIdx := len(digits) - 1

	return plus1(digits, lastIdx)
}

func plus1(digits []int, idx int) []int {
	if idx == -1 {
		return append([]int{1}, digits...)
	}
	digits[idx] = (digits[idx] + 1) % 10
	if digits[idx] == 0 {
		return plus1(digits, idx-1)
	}
	return digits
}