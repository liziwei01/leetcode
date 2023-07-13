/*
 * @Author: liziwei01
 * @Date: 2023-07-11 21:24:18
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-13 15:41:07
 * @Description: file content
 */
package sort

func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build maxheap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}

// siftDown implements the heap property on data[lo:hi].
// first is an offset into the array where the root of the heap lies.
func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}
