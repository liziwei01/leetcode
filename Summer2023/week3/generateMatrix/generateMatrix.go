/*
 * @Author: liziwei01
 * @Date: 2023-07-08 02:44:56
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 02:49:56
 * @Description: file content
 */
package generatematrix

func generateMatrix(n int) [][]int {
	columnLength := n
	tempColumnLength := columnLength
	rowLength := n
	tempRowLength := rowLength
	matrix := make([][]int, n, n)
	for i := 0; i != n; i++ {
		m := make([]int, n, n)
		matrix[i] = m
	}

	if n == 1 {
		return [][]int{{1}}
	}

	round := 0
	idx := 1
	for true {
		rightBegX := round
		rightBegY := rightBegX
		rightEndX := rightBegX
		rightEndY := rowLength - 1 - round
		downBegX := rightEndX
		downBegY := rightEndY
		downEndX := columnLength - 1 - round
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
				matrix[rightBegX][i] = idx
				idx++
			}
			break
		}
		if tempRowLength == 1 {
			for i := downBegX; i <= downEndX; i++ {
				matrix[i][downBegY] = idx
				idx++
			}
			break
		}

		for i := rightBegY; i != rightEndY; i++ {
			matrix[rightBegX][i] = idx
			idx++
		}
		for i := downBegX; i != downEndX; i++ {
			matrix[i][downBegY] = idx
			idx++
		}
		for i := leftBegY; i != leftEndY; i-- {
			matrix[leftBegX][i] = idx
			idx++
		}
		for i := upBegX; i != upEndX; i-- {
			matrix[i][upBegY] = idx
			idx++
		}
		round++
		tempRowLength = tempRowLength - 2
		tempColumnLength = tempColumnLength - 2
	}

	return matrix
}
