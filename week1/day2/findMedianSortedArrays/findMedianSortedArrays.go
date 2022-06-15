/*
 * @Author: liziwei01
 * @Date: 2022-06-14 11:53:19
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-14 14:23:49
 * @Description: file content
 */
package findmediansortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := joinTwoSortedArrays(nums1, nums2)
	if len(result)%2 == 0 {
		return float64(result[len(result)/2-1]+result[len(result)/2]) / 2
	}
	return float64(result[len(result)/2])
}

func joinTwoSortedArrays(nums1 []int, nums2 []int) []int {
	i := 0
	j := 0
	m := len(nums1)
	n := len(nums2)
	result := make([]int, 0)
	for i <= m && j <= n {
		if i == m {
			result = append(result, nums2[j:]...)
			break
		}
		if j == n {
			result = append(result, nums1[i:]...)
			break
		}
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	return result
}
