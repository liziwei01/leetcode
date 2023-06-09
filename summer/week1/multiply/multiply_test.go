/*
 * @Author: liziwei01
 * @Date: 2023-06-06 18:26:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-08 23:49:56
 * @Description: file content
 */
package multiply

import "testing"

func TestMultiply(t *testing.T) {
	i := New("237")
	j := New("284")

	i.Mul(j)
	t.Log(i.String())

	// i := New("6888")
	// j := New("49200")
	// i.Add(j)
	// t.Log(i.Str)
}
