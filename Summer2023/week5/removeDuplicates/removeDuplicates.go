/*
 * @Author: liziwei01
 * @Date: 2023-07-15 21:19:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-15 21:26:08
 * @Description: file content
 */
package removeduplicates

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	duplicates := 1
	for i := 1; i != len(nums); i++ {
		if nums[i] == nums[i-1] {
			duplicates++
		} else {
			duplicates = 1
		}
		if duplicates >= 3 {
			nums = append(nums[:i], nums[i+1:]...)
			duplicates--
			i--
		}
	}

	return len(nums)
}
