/*
 * @Author: liziwei01
 * @Date: 2023-07-05 09:25:29
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-06 05:45:22
 * @Description: file content
 */
package mypow

import "math"

func pow(x float64, n float64) float64 {
	return math.Pow(x, n)
}

func myPow(x float64, n int) float64 {
	// 剪枝
	//	Pow(x, ±0) = 1 for any x
	//	Pow(1, y) = 1 for any y
	//	Pow(x, 1) = x for any x
	if n == 0 {
		return 1
	}
	if x == 0 || x == 1 {
		return x
	}
	if x == -1 {
		if isOdd(n) {
			return -1
		}
		return 1
	}
	// if (n < -1000000 && math.Abs(x) > 1) || (n > 1000000 && math.Abs(x) < 1 && math.Abs(x) > 0) {
	// 	return 0
	// }
	if n < 0 {
		return myPow(1/x, -n)
	}

	pow := n
	ret := x
	var rem float64 = 1
	pre := ret
	// n减半处理
	for pow > 1 {
		remainder := pow % 2
		pow = pow / 2
		ret = ret * ret
		if remainder == 1 {
			rem = rem * pre
		}
		// 如果不变了就return
		if pre == ret {
			return ret
		}
		pre = ret
	}
	ret = ret * rem

	return ret
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func isOdd(n int) bool {
	return n%2 != 0
}
