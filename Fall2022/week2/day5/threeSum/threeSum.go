/*
 * @Author: liziwei01
 * @Date: 2022-06-17 08:25:27
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-27 18:19:05
 * @Description: file content
 */
package threeSum

import "fmt"

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	nums = Sort(nums...)
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}

	if nums[0] == nums[length-1] {
		if nums[0] == 0 {
			return [][]int{{0, 0, 0}}
		}
		return [][]int{}
	}

	i := 0
	j := 1
	k := length - 1

	for i < length-2 {
		if nums[i] > 0 {
			break
		}
		if i+3 < length-1 && nums[i] == nums[i+3] {
			i = i + 1
			j = i + 1
			k = length - 1
			continue
		}
		for j <= k {
			if j == k {
				i++
				j = i + 1
				k = length - 1
				break
			}
			if j+3 < k && nums[j] == nums[j+3] {
				j = j + 1
				k = length - 1
				continue
			}
			for k > j {
				if k-3 > j && nums[k] == nums[k-3] {
					k = k - 1
					continue
				}
				if nums[i]+nums[j]+nums[k] > 0 {
					k--
				} else if nums[i]+nums[j]+nums[k] == 0 {
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
					j++
					k = length - 1
					break
				} else {
					j++
					k = length - 1
					break
				}
			}
		}
	}
	return Deduplicate(ret)
}

func Sort(a ...int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func Deduplicate(nums [][]int) [][]int {
	m := make(map[string]bool)
	ret := [][]int{}

	for i := 0; i < len(nums); i++ {
		s := ""
		for j := 0; j < len(nums[i]); j++ {
			s += fmt.Sprint(nums[i][j])
		}
		if _, ok := m[s]; !ok {
			m[s] = true
			ret = append(ret, nums[i])
		}
	}

	return ret
}

// func Deduplicate(nums [][]int) [][]int {
// 	ret := [][]int{}
// 	length := len(nums)
// 	nums = Sortb(nums)
// 	for i := 0; i < length; i++ {
// 		if i+1 != length && nums[i][0] == nums[i+1][0] && nums[i][1] == nums[i+1][1] {
// 			continue
// 		}
// 		ret = append(ret, nums[i])
// 	}

// 	return ret
// }

// func Sortb(a [][]int) [][]int {
// 	for i := 0; i < len(a); i++ {
// 		for j := i + 1; j < len(a); j++ {
// 			if a[i][0] == a[j][0] && a[i][1] > a[j][1] {
// 				a[i], a[j] = a[j], a[i]
// 			}
// 		}
// 	}
// 	return a
// }
