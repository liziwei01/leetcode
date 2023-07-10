/*
 * @Author: liziwei01
 * @Date: 2023-07-08 15:06:26
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 15:17:45
 * @Description: file content
 */
package rotateright

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func len(l *ListNode) int {
	if l == nil {
		return 0
	}
	if l.Next != nil {
		return 1 + len(l.Next)
	}
	return 1
}

func rotateRight(head *ListNode, k int) *ListNode {
	n := len(head)
	if n == 0 {
		return head
	}
	k = k % n
	if k == 0 {
		return head
	}

	newTail := head
	for i := 0; i != n-1-k; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next

	tail := newTail
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = head
	newTail.Next = nil
	

	return newHead
}
