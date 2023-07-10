/*
 * @Author: liziwei01
 * @Date: 2023-07-06 18:25:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 01:26:52
 * @Description: file content
 */
package spiralorder

func spiralOrder(matrix [][]int) []int {
	columnLength := len(matrix)
	tempColumnLength := columnLength
	rowLength := len(matrix[0])
	tempRowLength := rowLength
	result := make([]int, 0)

	if columnLength == 1 {
		return matrix[0]
	}

	if rowLength == 1 {
		for i := 0; i != columnLength; i++ {
			result = append(result, matrix[i]...)
		}
		return result
	}

	n := 0
	for true {
		rightBegX := n
		rightBegY := rightBegX
		rightEndX := rightBegX
		rightEndY := rowLength - 1 - n
		downBegX := rightEndX
		downBegY := rightEndY
		downEndX := columnLength - 1 - n
		downEndY := downBegY
		leftBegX := downEndX
		leftBegY := downEndY
		leftEndX := leftBegX
		leftEndY := rightBegY
		upBegX := leftEndX
		upBegY := leftEndY
		upEndX := rightBegX
		// upEndY := rightBegY
		if tempColumnLength == 0 || tempRowLength == 0 {
			break
		}

		if tempColumnLength == 1 {
			for i := rightBegY; i <= rightEndY; i++ {
				result = append(result, matrix[rightBegX][i])
			}
			break
		}
		if tempRowLength == 1 {
			for i := downBegX; i <= downEndX; i++ {
				result = append(result, matrix[i][downBegY])
			}
			break
		}

		for i := rightBegY; i != rightEndY; i++ {
			result = append(result, matrix[rightBegX][i])
		}
		for i := downBegX; i != downEndX; i++ {
			result = append(result, matrix[i][downBegY])
		}
		for i := leftBegY; i != leftEndY; i-- {
			result = append(result, matrix[leftBegX][i])
		}
		for i := upBegX; i != upEndX; i-- {
			result = append(result, matrix[i][upBegY])
		}
		n++
		tempRowLength = tempRowLength - 2
		tempColumnLength = tempColumnLength - 2
	}

	return result
}
