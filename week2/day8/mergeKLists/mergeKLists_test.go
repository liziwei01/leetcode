/*
 * @Author: liziwei01
 * @Date: 2022-06-30 04:09:43
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-30 04:13:13
 * @Description: file content
 */
package mergeKLists

import "testing"

func TestMergeKLists(t *testing.T) {
	t.Log(mergeKLists([]*ListNode{createList([]int{2}), createList([]int{}), createList([]int{-1}) }))
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	pointer := head
	for i := 1; i < len(nums); i++ {
		pointer.Next = &ListNode{nums[i], nil}
		pointer = pointer.Next
	}
	return head
}
