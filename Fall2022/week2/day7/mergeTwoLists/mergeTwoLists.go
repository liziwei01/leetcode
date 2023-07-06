/*
 * @Author: liziwei01
 * @Date: 2022-06-29 01:55:56
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-29 02:12:09
 * @Description: file content
 */
package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	newList := list1
	if list1.Val > list2.Val {
		newList = list2
		list2 = list2.Next
	} else {
		list1 = list1.Next
	}
	head := newList

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			newList.Next = list2
			newList = newList.Next
			list2 = list2.Next
		} else {
			newList.Next = list1
			newList = newList.Next
			list1 = list1.Next
		}
	}

	if list1 != nil {
		newList.Next = list1
	}
	if list2 != nil {
		newList.Next = list2
	}

	return head
}
