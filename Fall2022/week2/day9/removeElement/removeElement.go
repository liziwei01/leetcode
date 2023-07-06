/*
 * @Author: liziwei01
 * @Date: 2022-07-02 04:20:56
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 04:20:56
 * @Description: file content
 */
package removeElement

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}