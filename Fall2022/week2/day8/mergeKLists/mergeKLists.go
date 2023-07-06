/*
 * @Author: liziwei01
 * @Date: 2022-06-30 04:09:38
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-30 04:14:48
 * @Description: file content
 */
package mergeKLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	for i := 0; i < length-1; i++ {
		lists[i+1] = mergeTwoLists(lists[i], lists[i+1])
	}

	return lists[length-1]
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
