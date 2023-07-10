/*
 * @Author: liziwei01
 * @Date: 2023-07-09 15:29:32
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-09 15:37:42
 * @Description: file content
 */
package mysqrt

func mySqrt(x int) int {
	y := x
	a, b := y/2, y
	for a*a > x {
		a /= 2
		b /= 2
	}
	for i := a; i != b+1; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return b
}
