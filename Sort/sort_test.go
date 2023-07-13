/*
 * @Author: liziwei01
 * @Date: 2023-07-12 22:22:42
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-13 17:52:51
 * @Description: file content
 */
package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	s := IntSlice{2, 1, 3, 5, 7, 6, 1, 2, 5, 6, 213, 2, 3, 4, 5, 6, 4}
	shellSort(s, 0, 4)
	t.Log(s)
}
