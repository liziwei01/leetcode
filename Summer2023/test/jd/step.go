/*
 * @Author: liziwei01
 * @Date: 2023-08-19 00:01:52
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-19 00:04:02
 * @Description: file content
 */
package jd

import (
	"fmt"
)

func step() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	board := make([]string, 0)
	for i := 0; i != n; i++ {
		str := ""
		fmt.Scan(&str)
		board = append(board, str)
	}

	stepBoard := make([][]int, 0)
	for i := 0; i != n; i++ {
		temp := make([]int, m)
		stepBoard = append(stepBoard, temp)
	}

	blocked := false
	for j := m - 1; j >= 0; j-- {
		if board[n-1][j] == '*' || blocked {
			stepBoard[n-1][j] = -1
			blocked = true
		} else {
			stepBoard[n-1][j] = m - j - 1
		}
	}

	blocked = false
	for i := n - 1; i >= 0; i-- {
		if board[i][m-1] == '*' || blocked {
			stepBoard[i][m-1] = -1
			blocked = true
		} else {
			stepBoard[i][m-1] = n - i - 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			if board[i][j] == '*' {
				stepBoard[i][j] = -1
			} else {
				if stepBoard[i+1][j+1] != -1 {
					stepBoard[i][j] = stepBoard[i+1][j+1] + 1
				} else {
					// cannot go directly
					if stepBoard[i+1][j] == -1 && stepBoard[i][j+1] == -1 {
						stepBoard[i][j] = -1
					} else if stepBoard[i+1][j] == -1 && stepBoard[i][j+1] != -1 {
						stepBoard[i][j] = stepBoard[i][j+1] + 1
					} else if stepBoard[i+1][j] != -1 && stepBoard[i][j+1] == -1 {
						stepBoard[i][j] = stepBoard[i+1][j] + 1
					} else {
						stepBoard[i][j] = min(stepBoard[i+1][j], stepBoard[i][j+1]) + 1
					}
				}
			}
		}
	}

	fmt.Printf("%d\n", stepBoard[0][0])
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
