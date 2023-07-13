/*
 * @Author: liziwei01
 * @Date: 2023-07-11 20:43:52
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-11 21:14:42
 * @Description: file content
 */
package sortcolors

func sortColors(nums []int) {
	n := len(nums)
	zeros := 0
	ones := 0
	for i := 0; i != n; i++ {
		if nums[i] == 0 {
			zeros++
		} else if nums[i] == 1 {
			ones++
		}
	}
	for i := 0; i != zeros; i++ {
		nums[i] = 0
	}
	for i := zeros; i != zeros+ones; i++ {
		nums[i] = 1
	}
	for i := zeros + ones; i != n; i++ {
		nums[i] = 2
	}
}
