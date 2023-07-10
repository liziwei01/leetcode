/*
 * @Author: liziwei01
 * @Date: 2023-07-06 16:04:24
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-06 18:22:26
 * @Description: file content
 */
package maxsubarray

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0

	for i := 0; i != len(nums); i++ {
		if sum <= 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		res = max(res, sum)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
