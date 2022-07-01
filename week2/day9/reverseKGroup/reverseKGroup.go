/*
 * @Author: liziwei01
 * @Date: 2022-07-01 18:11:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-01 19:30:17
 * @Description: file content
 */
package reverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}

	if k == 2 {
		return swapPairs(head)
	}

	beforeTail := head
	for i := 0; i != k-2; i++ {
		if beforeTail.Next != nil && beforeTail.Next.Next != nil {
			beforeTail = beforeTail.Next
		} else {
			return head
		}
	}

	// newHead := head.Next
	// newTail := head
	newHead := beforeTail.Next
	newTail := head
	afterHead := head.Next

	// newTail.Next = head.Next.Next
	// newHead.Next = head
	newTail.Next = beforeTail.Next.Next
	beforeTail.Next = head
	newHead.Next = afterHead

	newHead.Next = innerReverseKGroup(newHead.Next, k-2)
	newTail.Next = reverseKGroup(newTail.Next, k)

	return newHead
}

func innerReverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}

	if k == 2 {
		return innerSwapPairs(head)
	}

	beforeTail := head
	for i := 0; i != k-2; i++ {
		if beforeTail.Next != nil && beforeTail.Next.Next != nil {
			beforeTail = beforeTail.Next
		} else {
			return head
		}
	}

	// newHead := head.Next
	// newTail := head
	newHead := beforeTail.Next
	newTail := head
	afterHead := head.Next

	// newTail.Next = head.Next.Next
	// newHead.Next = head
	newTail.Next = beforeTail.Next.Next
	beforeTail.Next = head
	newHead.Next = afterHead

	newHead.Next = innerReverseKGroup(newHead.Next, k-2)

	return newHead
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

func innerSwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	newTail := head

	newTail.Next = head.Next.Next
	newHead.Next = head

	return newHead
}