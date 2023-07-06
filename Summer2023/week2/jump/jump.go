/*
 * @Author: liziwei01
 * @Date: 2023-06-26 11:32:42
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-26 12:05:58
 * @Description: file content
 */
package jump

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	return jumpTo(nums, 0)
}

func jumpTo(nums []int, idx int) int {
	if idx >= len(nums)-1 {
		return 0
	}

	if idx + nums[idx] >= len(nums)-1 {
		return 1
	}

	nextIdx := idx + 1
	farestTarget := nextIdx + nums[nextIdx]
	for i := nextIdx + 1; i <= idx + nums[idx]; i++ {
		target := i + nums[i]
		if farestTarget < target {
			nextIdx = i
			farestTarget = target
		}
	}
	
	return 1 + jumpTo(nums, nextIdx)
}