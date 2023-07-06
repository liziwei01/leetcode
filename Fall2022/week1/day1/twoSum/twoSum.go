/*
 * @Author: liziwei01
 * @Date: 2022-06-13 09:15:39
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-16 07:39:02
 * @Description: file content
 */
package twoSum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	ret := []int{}

	if len(nums) < 2 {
		return []int{}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				m[i] = 1
				m[j] = 1
			}
		}
	}

	for k := range m {
		if m[k] == 1 {
			ret = append(ret, k)
		}
	}

	return ret
}
