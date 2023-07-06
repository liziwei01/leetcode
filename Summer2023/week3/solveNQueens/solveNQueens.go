/*
 * @Author: liziwei01
 * @Date: 2023-07-06 08:46:45
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-06 09:10:04
 * @Description: file content
 */
package solvenqueens

func solveNQueens(n int) [][]string {
	chessBoards := make([][]string, 0)

	simpleboards := genAllSimpleboards(n)
	trueSimpleboards := make([][]int, 0)

	for i := 0; i != len(simpleboards); i++ {
		if check(n, simpleboards[i]) {
			trueSimpleboards = append(trueSimpleboards, simpleboards[i])
		}
	}

	for i := 0; i != len(trueSimpleboards); i++ {
		chessBoards = append(chessBoards, genChessboard(n, trueSimpleboards[i]))
	}
	
	return chessBoards
}

func genAllSimpleboards(n int) [][]int {
	firstSimpleBoard := make([]int, 0)
	for i := 0; i != n; i++ {
		firstSimpleBoard = append(firstSimpleBoard, i)
	}

	return permute(firstSimpleBoard)
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	permutations := make([][]int, 0)
	for i := 0; i != len(nums); i++ {
		newNum := deepCopy(&nums)
		newNum = append(newNum[:i], newNum[i+1:]...)
		permutes := permute(newNum)
		for j := 0; j != len(permutes); j++ {
			permutes[j] = append([]int{nums[i]}, permutes[j]...)
		}
		permutations = append(permutations, permutes...)
	}
	return permutations
}

func deepCopy(tmp *[]int) []int {
	res := make([]int, len(*tmp))
	copy(res, *tmp)
	return res
}

func check(n int, simpleboard []int) bool {
	for i := 0; i != n; i++ {
		line := simpleboard[i]
		for j := i + 1; j < n; j++ {
			if line == simpleboard[j]-(j-i) || line == simpleboard[j]+(j-i) {
				return false
			}
		}
	}

	return true
}

func genChessboard(n int, simpleboard []int) []string {
	strings := make([]string, 0)
	for i := 0; i != n; i++ {
		str := genLine(n, simpleboard[i])
		strings = append(strings, str)
	}

	return strings
}

func genLine(n int, idx int) string {
	str := ""
	for i := 0; i != n; i++ {
		if i == idx {
			str = str + "Q"
		} else {
			str = str + "."
		}
	}
	return str
}
