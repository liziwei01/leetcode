/*
 * @Author: liziwei01
 * @Date: 2022-08-25 09:14:29
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-25 10:18:12
 * @Description: file content
 */
package combinationSum

func combinationSum(candidates []int, target int) [][]int {
	if target < 1 {
		return nil
	}

	res := [][]int{}
	candidates = Sort(candidates...)
	length := len(candidates)

	// 树状结构
	for i := 0; i != length; i++ {
		// 因为sort过，后面肯定更大，不继续找了
		if candidates[i] > target {
			return res
		}
		if candidates[i] == target {
			res = append(res, []int{target})
			return res
		} else {
			// 不要走回头路，不能回去找前面的了
			ret := combinationSum(candidates[i:], target-candidates[i])
			if len(ret) == 0 {
				continue
			}

			for j := 0; j != len(ret); j++ {
				ret[j] = append(ret[j], candidates[i])
			}
			res = append(res, ret...)
		}
	}

	return res
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
