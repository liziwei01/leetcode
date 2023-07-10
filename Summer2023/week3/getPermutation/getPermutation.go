/*
 * @Author: liziwei01
 * @Date: 2023-07-08 02:52:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 03:17:11
 * @Description: file content
 */
package getpermutation

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	nums := []string{}
	for i := 1; i != n+1; i++ {
		nums = append(nums, strconv.Itoa(i))
	}
	return getPermutationRecursive(n, k, nums)
}

func getPermutationRecursive(n int, k int, nums []string) string {
	if n == 0 {
		return ""
	}
	// n = 3
	permuteN := permuteNum(n)
	// permuteN = 6
	lenofGroup := permuteN / n
	// lenofGroup = 2
	// numofGroups := n
	// k = 3
	groupIdx := k / lenofGroup
	k = k % lenofGroup
	// k = 1
	// 余数为0说明属于上一组末尾
	if k == 0 {
		groupIdx--
		k = lenofGroup
	}
	// groupIdx = 1
	// nums 中 groupIdx 所在的数就是string的第一个数
	newNums := deepCopy(&nums)
	return nums[groupIdx] + getPermutationRecursive(n-1, k, append(newNums[0:groupIdx], newNums[groupIdx+1:]...))
}

func permuteNum(n int) int {
	ret := 1
	for n != 0 {
		ret = ret * n
		n--
	}

	return ret
}

func deepCopy(tmp *[]string) []string {
	res := make([]string, len(*tmp))
	copy(res, *tmp)
	return res
}
