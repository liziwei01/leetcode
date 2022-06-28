/*
 * @Author: liziwei01
 * @Date: 2022-06-28 15:24:42
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-28 15:24:43
 * @Description: file content
 */
package fourSum

import "testing"

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fourSum(nums, target)
}
