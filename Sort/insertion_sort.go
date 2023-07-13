/*
 * @Author: liziwei01
 * @Date: 2023-07-11 21:23:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-13 03:20:43
 * @Description: file content
 */
package sort

// In insertion sort, the input array is also divided into two parts: a sorted part and an unsorted part.
// The algorithm picks up an element from the unsorted part and places it in the correct position in the sorted part, shifting the larger elements one position to the right.
// Insertion sort has a time complexity of O(n^2) in the worst case, but can perform better on partially sorted arrays, with a best-case time complexity of O(n).
// STABLE

// Main differences between selection sort and insertion sort:
// Selection sort scans the unsorted part to find the minimum element, while insertion sort scans the sorted part to find the correct position to place the element.
// Selection sort requires fewer swaps than insertion sort, but more comparisons.
// Insertion sort is more efficient than selection sort when the input array is partially sorted or almost sorted, while selection sort performs better when the array is highly unsorted.
// In summary, both algorithms have a similar time complexity, but their selection and placement methods differ. The choice between them depends on the characteristics of the input data and the specific requirements of the problem at hand.

// insertionSort sorts data[a:b] using insertion sort.
func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}