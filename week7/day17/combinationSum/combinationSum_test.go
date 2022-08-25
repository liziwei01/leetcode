/*
 * @Author: liziwei01
 * @Date: 2022-08-25 10:07:54
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-25 10:08:56
 * @Description: file content
 */
package combinationSum

import "testing"

func TestCombinationSum(t *testing.T) {
	ret := combinationSum([]int{2, 3, 6, 7}, 7)

	t.Log(ret)
}
