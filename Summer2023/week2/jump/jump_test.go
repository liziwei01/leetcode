/*
 * @Author: liziwei01
 * @Date: 2023-06-26 11:46:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-26 12:03:12
 * @Description: file content
 */
package jump

import "testing"

func TestJump(t *testing.T) {
	in := []int{4,1,1,3,1,1,1}
	t.Log(jump(in))
}
