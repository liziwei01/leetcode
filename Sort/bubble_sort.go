/*
 * @Author: liziwei01
 * @Date: 2023-07-13 03:17:43
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-13 03:19:52
 * @Description: file content
 */
package sort

// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 针对所有的元素重复以上的步骤，除了最后一个。
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
// STABLE
func bubbleSort(data Interface, a, b int) {
	length := b - a
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if !data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}
