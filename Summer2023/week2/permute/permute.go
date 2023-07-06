/*
 * @Author: liziwei01
 * @Date: 2023-06-28 10:14:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-28 10:40:05
 * @Description: file content
 */
package permute

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