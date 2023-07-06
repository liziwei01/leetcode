/*
 * @Author: liziwei01
 * @Date: 2022-06-16 07:23:59
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-16 07:39:10
 * @Description: file content
 */
package maxArea

func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		area := Area(i, j, Min(height[i], height[j]))
		max = Max(max, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}

func Max(a ...int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(a ...int) int {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func Area(xi, xj, height int) int {
	return Abs(xi-xj) * height
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
