/*
 * @Author: liziwei01
 * @Date: 2023-07-08 01:36:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 01:55:58
 * @Description: file content
 */
package merge

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	return mergeRecursive(intervals)
}

func mergeRecursive(intervals [][]int) [][]int {
	res := [][]int{}
	done := true
	// 单独确认最后一个元素是否处理过
	lastDone := false
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][len(intervals[i])-1] >= intervals[i+1][0] {
			res = append(res, []int{min(intervals[i][0], intervals[i+1][0]), max(intervals[i][len(intervals[i])-1], intervals[i+1][len(intervals[i+1])-1])})
			if i == len(intervals)-2 {
				lastDone = true
			}
			i++
			done = false
		} else {
			res = append(res, intervals[i])
		}
	}
	if done {
		return intervals
	}
	if !lastDone {
		res = append(res, intervals[len(intervals)-1])
	}

	return mergeRecursive(res)
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}