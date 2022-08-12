/*
 * @Author: liziwei01
 * @Date: 2022-08-12 17:01:08
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-12 18:34:46
 * @Description: file content
 */
package search

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	// 寻找头的位置
	head := findHead(nums)
	// 如果头就在头上，说明全部是升序，直接查找
	if head == 0 {
		return binarySearch(nums, target)
	}
	// 如果不是0
	idx := binarySearch(nums[:head], target)
	if idx != -1 {
		return idx
	}

	idx = binarySearch(nums[head:], target)
	if idx != -1 {
		return head + idx
	}

	return -1
}

// 找原来的头在哪里
func findHead(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	// 1 寻找是在哪里旋转的
	// 有以下几种可能
	// 0. 全部是倒序（只有两个元素时可能发生）
	// 1. 全部是升序，旋转在第一个位置
	// 2. 旋转在中间，旋转的位置在中间元素的左边或正好是中间元素。2转：12345 - > 23451，3转：12345 -> 34512，则现在的中间元素比头大，比尾大
	// 3. 旋转在中间，旋转的位置在中间元素的右边。4转：12345 - > 45123则现在的中间元素比头小，比尾小
	// 中间值
	mid := nums[len(nums)/2]
	// 0.
	if mid == nums[len(nums)-1] {
		// 说明正好只有俩元素而且倒序了
		return 1
	}
	// 1.
	if nums[0] < mid && mid < nums[len(nums)-1] {
		return 0
	}

	// 2. 如果中间值比头大，说明旋转的位置在中间元素的左边 (转后，头到了后面)
	if mid > nums[0] {
		return len(nums)/2 + 1 + findHead(nums[len(nums)/2+1:])
	}

	// 3. 如果中间值比头小，说明旋转的位置在中间元素的右边 (转后，头到了前面或者正中央)
	return findHead(nums[:len(nums)/2+1])
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
	return -1
}
