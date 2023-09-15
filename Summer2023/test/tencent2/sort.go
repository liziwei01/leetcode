/*
 * @Author: liziwei01
 * @Date: 2023-09-15 08:01:02
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-15 08:11:40
 * @Description: file content
 */
package tencent2

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 返回一个指向链表头部的指针。
 * @param a ListNode类一维数组 指向这些数链的开头
 * @return ListNode类
 */
func solve(a []*ListNode) *ListNode {
	// write code here
	nodes := make([]int, 0)
	for i := 0; i < len(a); i++ {
		b := a[i]
		for b != nil {
			nodes = append(nodes, b.Val)
			b = b.Next
		}
	}

	sort.Ints(nodes)

	head := &ListNode{
		Val: nodes[0],
		Next: nil,
	}
	tmp := head

	for i := 1; i < len(nodes); i++ {
		tmp.Next = &ListNode{
			Val: nodes[i],
			Next: nil,
		}
		tmp = tmp.Next
	}

	return head
}
