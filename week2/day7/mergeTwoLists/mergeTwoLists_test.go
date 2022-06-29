/*
 * @Author: liziwei01
 * @Date: 2022-06-29 01:56:02
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-29 02:10:22
 * @Description: file content
 */
package mergeTwoLists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	t.Log(mergeTwoLists(createList([]int{2}), createList([]int{1})))
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
