/*
 * @Author: liziwei01
 * @Date: 2023-06-04 14:16:16
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-11 09:47:55
 * @Description: file content
 */
package combinationSum2

import (
	"sort"
)

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

func combinationSum2(candidates []int, target int) [][]int {
	// quick sort
	sort.IntSlice(candidates).Sort()
	// recursive
	res := combSum2(candidates, target)

	// 去重
	return uniqueSlice(res)
}

func combSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)

	// 剪枝
	for i := 0; i != len(candidates); i++ {
		if candidates[i] > target {
			candidates = candidates[0:i]
			break
		}
	}

	// recursive
	for i := 0; i != len(candidates); i++ {
		// 减去当前数
		newTarget := target - candidates[i]
		if newTarget == 0 {
			// 如果相等，当前数即为一个解
			res = append(res, []int{target})
		} else {
			// 如果不相等，继续recursive直到找到相等或减空
			newCandidates := deepCopy(&candidates)
			newCandidates = newCandidates[i+1:]
			ret := combSum2(newCandidates, newTarget)
			// 将所有找到的解加上减去的那个数字
			if len(ret) != 0 {
				ret = allAppend(ret, candidates[i])
				res = append(res, ret...)
			}
		}
	}

	return res
}

func allAppend(candidates [][]int, target int) [][]int {
	for i := 0; i != len(candidates); i++ {
		candidates[i] = append(candidates[i], target)
	}

	return candidates
}

func deepCopy(tmp *[]int) []int {
	res := make([]int, len(*tmp))
	copy(res, *tmp)
	return res
}

func uniqueSlice(s [][]int) [][]int {
	if len(s) <= 1 {
		return s
	}

	ret := make([][]int, 0)

	sort.Sort(IntIntSlice(s))

	for i := 0; i != len(s) - 1; i++ {
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
