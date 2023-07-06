/*
 * @Author: liziwei01
 * @Date: 2023-06-28 10:47:21
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-28 10:48:29
 * @Description: file content
 */
package permuteunique

import "sort"

func permuteUnique(nums []int) [][]int {
	permutations := permute(nums)

	// 去重
	return uniqueSlice(permutations)
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	permutations := make([][]int, 0)
	for i := 0; i != len(nums); i++ {
		newNum := deepCopy(&nums)
		newNum = append(newNum[:i], newNum[i+1:]...)
		permutes := permute(newNum)
		for j := 0; j != len(permutes); j++ {
			permutes[j] = append([]int{nums[i]}, permutes[j]...)
		}
		permutations = append(permutations, permutes...)
	}
	return permutations
}

func deepCopy(tmp *[]int) []int {
	res := make([]int, len(*tmp))
	copy(res, *tmp)
	return res
}


type IntIntSlice [][]int

func (s IntIntSlice) Len() int {
	return len(s)
}

func (s IntIntSlice) Less(i, j int) bool {
	k := 0
	for k < len(s[i]) && k < len(s[j]) && s[i][k] == s[j][k] {
		k++
	}

	if !(k < len(s[i]) && k < len(s[j])) {
		return true
	}

	return s[i][k] < s[j][k]
}

func (s IntIntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func uniqueSlice(s [][]int) [][]int {
	if len(s) <= 1 {
		return s
	}

	ret := make([][]int, 0)

	sort.Sort(IntIntSlice(s))

	for i := 0; i != len(s)-1; i++ {
		if !intSliceEqual(s, i, i+1) {
			ret = append(ret, s[i])
		}
	}

	return append(ret, s[len(s)-1])
}

func intSliceEqual(s [][]int, i, j int) bool {
	k := 0
	for k < len(s[i]) && k < len(s[j]) && s[i][k] == s[j][k] {
		k++
	}

	if k >= len(s[i]) && k >= len(s[j]) {
		return true
	}

	return false
}
