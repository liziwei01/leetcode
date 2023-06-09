/*
 * @Author: liziwei01
 * @Date: 2023-06-04 20:47:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-04 22:09:47
 * @Description: file content
 */
package combinationSum2

import (
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	arr := []int{10,1,2,7,6,1,5}

	can := combinationSum2(arr, 8)

	t.Log(can)
}
