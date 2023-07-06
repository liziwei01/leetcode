/*
 * @Author: liziwei01
 * @Date: 2022-07-02 04:17:24
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-23 21:36:56
 * @Description: file content
 */
package removeDuplicates

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}