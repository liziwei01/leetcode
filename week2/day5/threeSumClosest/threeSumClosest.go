/*
 * @Author: liziwei01
 * @Date: 2022-06-27 18:22:16
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-27 18:36:02
 * @Description: file content
 */
package threeSumClosest

func threeSumClosest(nums []int, target int) int {
	ret := 0
	nums = Sort(nums...)
	length := len(nums)

	if length <= 3 {
		for _, v := range nums {
			ret += v
		}

		return ret
	}

	i := 0
	j := 1
	k := length - 1
	ret = nums[0] + nums[1] + nums[2]

	for i < length-2 {
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
				sum := nums[i] + nums[j] + nums[k]
				if sum > target {
					if Abs(sum-target) < Abs(ret-target) {
						ret = sum
					}
					k--
				} else if sum == target {
					return target
				} else {
					if Abs(sum-target) < Abs(ret-target) {
						ret = sum
					}
					j++
					k = length - 1
					break
				}
			}
		}
	}
	return ret
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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
