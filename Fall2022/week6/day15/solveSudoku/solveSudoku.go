/*
 * @Author: liziwei01
 * @Date: 2022-08-16 19:47:15
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-17 02:05:03
 * @Description: file content
 */
package solveSudoku

func solveSudoku(board [][]byte) {
	candidates := make([][][]byte, 9, 9)

	for !solvedSudoku(board) {
		priviousBoard := deepCopyBoard(board)
		// 1 把100%能填的候选人填上
		hasCandidate := trySolveSudoku(board, candidates)
		// 如果有至少一个空格子没有候选人，说明这个分支走到尽头了，直接剪枝
		if !hasCandidate {
			return
		}
		// 2 没有任何一个空格还能直接填了，开始使用3x3其他格子候选人的排除法
		if equalBoard(priviousBoard, board) {
			fillinUniqueCandidateinblock(board, candidates)
		}
		// 3 如果还是找不到，没办法了，只能硬猜了
		if equalBoard(priviousBoard, board) {
			solved := guessBoard(board)
			if solved {
				return
			}
		}
		if !isValidSudoku(board) {
			return
		}
	}

	return
}

func guessBoard(guessedBoard [][]byte) bool {
	guessedCandidates := getAllCandidates(guessedBoard)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if guessedBoard[i][j] == '.' {
				if len(guessedCandidates[i][j]) == 0 {
					return false
				}
				// 一个一个候选人的试
				guessedCandidatesij := guessedCandidates[i][j]
				// guessedCandidatesij := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
				for _, candidate := range guessedCandidatesij {
					guessedBoard[i][j] = candidate
					if isValidSudoku(guessedBoard) {
						if solvedSudoku(guessedBoard) {
							return true
						}
						if guessBoard(guessedBoard) {
							return true
						}
					} else {
						return false
					}
					guessedBoard[i][j] = '.'
				}

				return false
			}
		}
	}

	return true
}

func equalBoard(board1, board2 [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board1[i][j] != board2[i][j] {
				return false
			}
		}
	}
	return true
}

func solvedSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return false
			}
		}
	}
	return true
}

// 走一遍数独，候选人只有一个的时候，直接填入
func trySolveSudoku(board [][]byte, candidates [][][]byte) bool {
	for i := 0; i < 9; i++ {
		candidates[i] = make([][]byte, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				candidates[i][j] = getCandidates(board, i, j)
				if len(candidates[i][j]) == 0 {
					return false
				}
				if len(candidates[i][j]) == 1 {
					board[i][j] = candidates[i][j][0]
				}
			}
		}
	}

	return true
}

// 假如某个小格子的某个候选人和3x3内以及横竖的其他格子的任何候选人都没有重合，那么这个小格子的候选人就只有一个了
func fillinUniqueCandidateinblock(board [][]byte, candidates [][][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for _, candidate := range candidates[i][j] {
					otherCandidates := getBlockOtherCandidate(board, candidates, i, j)
					if uniqueCandidateinBlock(candidate, otherCandidates) {
						board[i][j] = candidate
						candidates = getAllCandidates(board)
						break
					}
				}
			}
		}
	}
}

func uniqueCandidateinBlock(candidate byte, otherCandidate []byte) bool {
	for _, other := range otherCandidate {
		if other == candidate {
			return false
		}
	}
	return true
}

func getAllCandidates(board [][]byte) [][][]byte {
	usedBoard := deepCopyBoard(board)
	candidates := make([][][]byte, 9, 9)
	for i := 0; i < 9; i++ {
		candidates[i] = make([][]byte, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				candidates[i][j] = getCandidates(usedBoard, i, j)
			}
		}
	}

	return candidates
}

func getCandidates(board [][]byte, i, j int) []byte {
	candidates := make([]byte, 0)
	for k := 0; k < 9; k++ {
		board[i][j] = byte(k + '1')
		if isValidSudoku(board) {
			candidates = append(candidates, byte(k+'1'))
		}
		board[i][j] = '.'
	}
	return candidates
}

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

// 获取除了自己以外，3x3内以及横竖的其他格子的所有候选人
func getBlockOtherCandidate(board [][]byte, candidates [][][]byte, i, j int) []byte {
	// 先3x3内的候选人
	blockRow := i / 3
	blockCol := j / 3
	otherCandidates := make([]byte, 0)
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			m := blockRow*3 + k
			n := blockCol*3 + l
			if m != i || n != j {
				for _, candidate := range candidates[m][n] {
					if candidate != '.' {
						otherCandidates = append(otherCandidates, candidate)
					}
				}
			}
		}
	}

	// 横竖的候选人
	// for k := 0; k < 9; k++ {
	// 	if k != i {
	// 		for _, candidate := range candidates[k][j] {
	// 			if candidate != '.' {
	// 				otherCandidates = append(otherCandidates, candidate)
	// 			}
	// 		}
	// 	}
	// }
	// for k := 0; k < 9; k++ {
	// 	if k != j {
	// 		for _, candidate := range candidates[i][k] {
	// 			if candidate != '.' {
	// 				otherCandidates = append(otherCandidates, candidate)
	// 			}
	// 		}
	// 	}
	// }

	return otherCandidates
}

func deepCopyBoard(board [][]byte) [][]byte {
	newBoard := make([][]byte, 9, 9)
	for i := 0; i < 9; i++ {
		newBoard[i] = make([]byte, 9, 9)
		for j := 0; j < 9; j++ {
			newBoard[i][j] = board[i][j]
		}
	}
	return newBoard
}
