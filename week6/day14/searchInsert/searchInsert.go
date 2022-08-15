/*
 * @Author: liziwei01
 * @Date: 2022-08-14 18:46:21
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-14 18:49:04
 * @Description: file content
 */
package searchInsert

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target)
}

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high + 1
}
