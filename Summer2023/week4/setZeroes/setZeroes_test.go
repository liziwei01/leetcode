/*
 * @Author: liziwei01
 * @Date: 2023-07-11 01:06:55
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-11 01:06:55
 * @Description: file content
 */
package setzeroes

import "testing"

func TestXxx(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	t.Log(matrix)
}
