/*
 * @Author: liziwei01
 * @Date: 2022-06-30 04:17:26
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-30 04:23:59
 * @Description: file content
 */
package swapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	newTail := head

	newTail.Next = head.Next.Next
	newHead.Next = head

	newTail.Next = swapPairs(newTail.Next)

	return newHead
}
