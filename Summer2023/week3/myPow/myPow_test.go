/*
 * @Author: liziwei01
 * @Date: 2023-07-05 15:43:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-06 05:42:52
 * @Description: file content
 */
package mypow

import "testing"

func TestMyPow(t *testing.T) {
	t.Log(myPow(1.0000000000001, -2147483648))
	// t.Log(myPow(2, 12))
}
