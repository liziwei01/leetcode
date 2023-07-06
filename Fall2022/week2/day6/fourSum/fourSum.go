/*
 * @Author: liziwei01
 * @Date: 2022-06-28 15:11:55
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-28 15:34:20
 * @Description: file content
 */
package fourSum

import "fmt"

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	nums = Sort(nums...)
	length := len(nums)

	if length < 4 {
		return [][]int{}
	}

	if nums[0] == nums[length-1] {
		if nums[0] == 0 && target == 0 {
			return [][]int{{0, 0, 0, 0}}
		}
	}

	i := 0
	j := 1
	k := 2
	l := length - 1

	for i < length-3 {
		if i+4 < length-1 && nums[i] == nums[i+4] {
			i++
			j = i + 1
			k = j + 1
			l = length - 1
			continue
		}
		for j <= k {
			if k >= length || j == k {
				i++
				j = i + 1
				k = j + 1
				l = length - 1
				break
			}
			if j+4 < k && nums[j] == nums[j+4] {
				j++
				k = j + 1
				l = length - 1
				continue
			}
			for k <= l {
				if k == l {
					j++
					k = j + 1
					l = length - 1
					break
				}
				if k+4 < l && nums[k] == nums[k+4] {
					k = k + 1
					l = length - 1
					continue
				}
				for l > k {
					if l-4 > k && nums[l] == nums[l-4] {
						l = l - 1
						continue
					}
					sum := nums[i] + nums[j] + nums[k] + nums[l]
					if sum > target {
						l--
					} else if sum == target {
						ret = append(ret, []int{nums[i], nums[j], nums[k], nums[l]})
						k++
						l = length - 1
						break
					} else {
						k++
						l = length - 1
						break
					}
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
