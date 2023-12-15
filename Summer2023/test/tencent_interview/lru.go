/*
 * @Author: liziwei01
 * @Date: 2023-12-15 14:17:03
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-12-15 14:17:04
 * @Description: file content
 */
/*
 * @Author: liziwei01
 * @Date: 2023-09-13 16:56:04
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-12-15 14:14:33
 * @Description: file content
 */
package tencentinterview

import (
	"fmt"
)

// 运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

// 获取数据 get(key)：

//     如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1

// 写入数据 put(key, value)：

//     如果关键字已经存在，则变更其数据值

//     如果关键字不存在，则插入该组「关键字/值」

//     当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间

// 示例：

type ListNode struct {
	Key      int
	Val      int
	UsedTime int64
	Prev     *ListNode
	Next     *ListNode
}

type LRUCache struct {
	Root     *ListNode
	has      int
	capacity int
}

func (this *LRUCache) init(capacity int) {
	this.Root = nil
	this.has = 0
	this.capacity = capacity
}

func (this *LRUCache) get(key int) int {
	if this.Root == nil {
		return -1
	}

	th := this.Root
	for th != nil {
		if th.Key == key {
			// Move this node to the head of the list.
			if th.Prev != nil {
				th.Prev.Next = th.Next
			}
			if th.Next != nil {
				th.Next.Prev = th.Prev
			}
			th.Prev = nil
			th.Next = this.Root
			this.Root.Prev = th
			this.Root = th
			return th.Val
		}
		th = th.Next
	}
	return -1
}

func (this *LRUCache) put(key int, value int) {
	var now int64 = 0
	// 为空
	if this.Root == nil {
		this.Root = &ListNode{
			Key:      key,
			Val:      value,
			UsedTime: now,
			Prev:     nil,
			Next:     nil,
		}
		this.has++
		return
	}
	// 存在
	th := this.Root
	for th != nil {
		if th.Key == key {
			// 找到了，更新，return
			th.Val = value
			th.UsedTime = now
			// Move this node to the head of the list.
			if th.Prev != nil {
				th.Prev.Next = th.Next
			}
			if th.Next != nil {
				th.Next.Prev = th.Prev
			}
			th.Prev = nil
			th.Next = this.Root
			this.Root.Prev = th
			this.Root = th
			return
		}
		th = th.Next
	}
	// 未满不存在,加一个
	if this.has < this.capacity {
		this.has++
		t := this.Root
		this.Root = &ListNode{
			Key:      key,
			Val:      value,
			UsedTime: now,
			Prev:     nil,
			Next:     t,
		}
		t.Prev = this.Root
		return
	}
	// 已满，覆写尾部节点
	th = this.Root
	for th.Next != nil {
		th = th.Next
	}
	th.Key = key
	th.Val = value
	th.UsedTime = now

	// Move this node to the head of the list.
	if th.Prev != nil {
		th.Prev.Next = th.Next
	}
	if th.Next != nil {
		th.Next.Prev = th.Prev
	}
	th.Prev = nil
	th.Next = this.Root
	this.Root.Prev = th
	this.Root = th
}

func lru() {
	cache := LRUCache{}
	cache.init(2)
	cache.put(1, 1)
	cache.put(2, 2)
	v := cache.get(1) // 返回 1
	cache.put(3, 3)   // 该操作会使得关键字 2 作废
	v = cache.get(2)  // 返回 -1 (未找到)
	cache.put(4, 4)   // 该操作会使得关键字 1 作废
	v = cache.get(1)  // 返回 -1 (未找到)
	v = cache.get(3)  // 返回 3
	v = cache.get(4)  // 返回 4
	fmt.Println(v)
}
