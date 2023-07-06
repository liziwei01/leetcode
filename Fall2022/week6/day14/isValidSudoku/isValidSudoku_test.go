/*
 * @Author: liziwei01
 * @Date: 2022-08-14 18:55:48
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-14 18:58:05
 * @Description: file content
 */
package isValidSudoku

import "testing"

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	if !isValidSudoku(board) {
		t.Error("isValidSudoku failed")
	}
}
