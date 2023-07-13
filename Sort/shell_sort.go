/*
 * @Author: liziwei01
 * @Date: 2023-07-13 17:49:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-13 17:52:32
 * @Description: file content
 */
package sort

// shellSort sorts data[a:b] using shell sort.
func shellSort(data Interface, a, b int) {
	h := 1
	for h < (b-a)/2 {
		h = h*2 + 1
	}
	for ; h >= 1; h /= 2 {
		for i := a + h; i < b; i++ {
			for j := i; j >= h && data.Less(j, j-h); j = j - h {
				data.Swap(j, j-h)
			}
		}
	}
}
