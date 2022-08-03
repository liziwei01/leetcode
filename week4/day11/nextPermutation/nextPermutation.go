/*
 * @Author: liziwei01
 * @Date: 2022-08-03 02:52:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-03 03:14:42
 * @Description: file content
 */
package nextPermutation

import "sort"

func nextPermutation(nums []int) {
	numsLen := len(nums)
	if numsLen <= 1 {
		return
	}

	for i := numsLen - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			// 找一个所有比i-1大的数字里面最小的数字
			minIdx := i
			for j := i + 1; j != numsLen; j++ {
				if nums[j] > nums[i-1] && nums[j] < nums[minIdx] {
					minIdx = j
				}
			}
			// 交换i-1和minIdx
			nums[i-1], nums[minIdx] = nums[minIdx], nums[i-1]
			// i 后面的全部翻转
			sort.Ints(nums[i:])
			return
		}
	}

	// 已经是最大排列了，那就全部翻转做成最小排列
	sort.Ints(nums)
}
