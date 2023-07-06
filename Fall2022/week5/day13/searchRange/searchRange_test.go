/*
 * @Author: liziwei01
 * @Date: 2022-08-14 13:47:47
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-14 13:47:48
 * @Description: file content
 */
package searchRange

import "testing"

func TestSearchRange(t *testing.T) {
	nums := []int{1,2,3,3,3,3,4,5,9}
	target := 3
	t.Log(searchRange(nums, target))
}
