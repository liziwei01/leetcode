/*
 * @Author: liziwei01
 * @Date: 2022-08-14 13:33:26
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-14 13:56:14
 * @Description: file content
 */
package searchRange

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	mid, low, high := binarySearch(nums, target)
	if mid == -1 {
		return []int{-1, -1}
	}

	frontEdge := findEdge(nums, target, low, mid, true)
	backEdge := findEdge(nums, target, mid, high, false)

	return []int{frontEdge, backEdge}
}

func binarySearch(nums []int, target int) (int, int, int) {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid, low, high
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, -1, -1
}

func findEdge(nums []int, target, low, high int, findFrontEdge bool) int {
	if low == high {
		if nums[low] == target {
			return low
		}
	}

	// 由于底下除二取的是floor，所以low high只差一的时候，找尾巴会陷入死循环，特地加一个判断
	if !findFrontEdge {
		if low == high-1 {
			if nums[high] == target {
				return high
			} else {
				return low
			}
		}
	}

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			if findFrontEdge {
				return findEdge(nums, target, low, mid, findFrontEdge)
			} else {
				return findEdge(nums, target, mid, high, findFrontEdge)
			}
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
