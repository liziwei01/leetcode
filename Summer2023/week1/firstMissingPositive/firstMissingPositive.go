/*
 * @Author: liziwei01
 * @Date: 2023-06-05 21:19:32
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-06 00:46:53
 * @Description: file content
 */
package firstmissingpositive

func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i != length; i++ {
		if nums[i] <= 0 {
			nums[i] = length + 1
		}
	}

	for i := 0; i != length; i++ {
		nums_i := abs(nums[i])
			if nums_i <= length {
			nums[nums_i - 1] = -abs(nums[nums_i - 1])
		}
	}

	for i := 0; i != length; i++ {
		if nums[i] >= 0 {
			return i + 1
		}
	}

	return length + 1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}