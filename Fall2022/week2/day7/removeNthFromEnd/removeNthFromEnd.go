/*
 * @Author: liziwei01
 * @Date: 2022-06-29 01:21:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-29 01:35:56
 * @Description: file content
 */
package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	prePointer := head
	tail := head.Next
	length := 2
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	if length == n {
		return head.Next
	}

	for i := 0; i < length-n-1; i++ {
		prePointer = prePointer.Next
	}

	prePointer.Next = prePointer.Next.Next

	return head
}
