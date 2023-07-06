/*
 * @Author: liziwei01
 * @Date: 2022-06-29 01:29:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-29 01:33:50
 * @Description: file content
 */
package removeNthFromEnd

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	t.Log(removeNthFromEnd(createList([]int{1, 2}), 2))
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
