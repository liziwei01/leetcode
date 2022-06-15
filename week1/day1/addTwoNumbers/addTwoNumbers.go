/*
 * @Author: liziwei01
 * @Date: 2022-06-13 09:25:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-14 14:24:05
 * @Description: file content
 */
package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var ret *ListNode = &ListNode{}
	ret.Val = l1.Val + l2.Val

	if ret.Val < 10 {
		ret.Next = addTwoNumbers(l1.Next, l2.Next)
	} else {
		ret.Val = ret.Val - 10
		ret.Next = addTwoNumbers(l1.Next, l2.Next)
		ret.Next = addTwoNumbers(ret.Next, &ListNode{Val: 1})
	}

	return ret
}
