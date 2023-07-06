/*
 * @Author: liziwei01
 * @Date: 2022-07-01 18:44:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-01 19:23:20
 * @Description: file content
 */
package reverseKGroup

import "testing"

func TestReverseKGroup(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	head := createList(nums)
	result := reverseKGroup(head, 3)
	for result != nil {
		t.Log(result.Val)
		result = result.Next
	}
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
