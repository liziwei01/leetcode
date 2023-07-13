/*
 * @Author: liziwei01
 * @Date: 2023-07-11 21:24:50
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-13 03:20:54
 * @Description: file content
 */
package sort

// In selection sort, the input array is divided into two parts: a sorted part and an unsorted part.
// The algorithm repeatedly finds the minimum element in the unsorted part and swaps it with the leftmost element of the unsorted part, thus expanding the sorted part by one element.
// Selection sort has a time complexity of O(n^2) in all cases.
// NOT STABLE

// selectionSort sorts data[a:b] using selection sort.
func selectionSort(data Interface, a, b int) {
	for i := a; i < b; i++ {
		lo := i
		for j := i + 1; j < b; j++ {
			if data.Less(j, lo) {
				lo = j
			}
		}
		data.Swap(i, lo)
	}
}