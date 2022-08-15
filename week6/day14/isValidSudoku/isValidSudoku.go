/*
 * @Author: liziwei01
 * @Date: 2022-08-14 18:50:25
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-14 19:03:23
 * @Description: file content
 */
package isValidSudoku

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				row := checkRow(board, i, j)
				col := checkCol(board, i, j)
				block := checkBlock(board, i, j)
				if !row || !col || !block {
					return false
				}
			}
		}
	}
	return true
}

func checkRow(board [][]byte, i, j int) bool {
	for k := 0; k < 9; k++ {
		if k != j && board[i][k] == board[i][j] {
			return false
		}
	}
	return true
}

func checkCol(board [][]byte, i, j int) bool {
	for k := 0; k < 9; k++ {
		if k != i && board[k][j] == board[i][j] {
			return false
		}
	}
	return true
}

func checkBlock(board [][]byte, i, j int) bool {
	blockRow := i / 3
	blockCol := j / 3
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			m := blockRow*3 + k
			n := blockCol*3 + l
			if (m != i || n != j) && board[m][n] == board[i][j] {
				return false
			}
		}
	}
	return true
}
