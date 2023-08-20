/*
 * @Author: liziwei01
 * @Date: 2023-07-27 22:01:24
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-27 23:30:26
 * @Description: file content
 */
package largestRectangleArea

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	largestArea := 0

	for i := 0; i < len(heights); i++ {
		height := heights[i]
		for j := i; j < len(heights); j++ {
			length := j - i + 1
			if heights[j] < height {
				height = heights[j]
			}
			area := height * length
			if area > largestArea {
				largestArea = area
			}
		}
	}

	return largestArea
}

func LargestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	largestArea := 0
	a, b := 0, len(heights)

	for a < len(heights) {
		tmpArea := getTmpArea(heights, a, b)
		if tmpArea >= largestArea {
			largestArea = tmpArea
		}
		a++
	}

	return largestArea
}

func getTmpArea(heights []int, a, b int) int {
	length := b - a
	height := heights[a]

	for i := a + 1; i < b; i++ {
		if heights[i] < height {
			height = heights[i]
		}
	}

	return length * height
}
