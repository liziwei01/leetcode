/*
 * @Author: liziwei01
 * @Date: 2022-06-14 13:30:35
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-14 14:23:29
 * @Description: file content
 */
package intreverse

import (
	"fmt"
)

func reverse(x int) int {
	isMinus := false
	if x < 0 {
		isMinus = true
		x = -x
	}
	xAbsStr := fmt.Sprintf("%d", x)

	xAbsByte := []byte(xAbsStr)
	for i := 0; i < len(xAbsByte)/2; i++ {
		xAbsByte[i], xAbsByte[len(xAbsByte)-i-1] = xAbsByte[len(xAbsByte)-i-1], xAbsByte[i]
	}

	if ExceedRange(xAbsByte, isMinus) {
		return 0
	}

	if isMinus {
		return -1 * Int(xAbsByte)
	}
	return Int(xAbsByte)
}

func ExceedRange(xAbsByte []byte, isMinus bool) bool {
	length := len(xAbsByte)
	benchmark := []byte("2147483647")
	if length > 10 {
		return true
	}
	if length < 10 {
		return false
	}
	if isMinus {
		benchmark = []byte("2147483648")
	} else {
		benchmark = []byte("2147483647")
	}
	for i := 0; i < len(benchmark); i++ {
		x := xAbsByte[i]
		b := benchmark[i]
		if x > b {
			return true
		} else if x == b {
			continue
		} else {
			return false
		}
	}
	return false
}

func Int(b []byte) int {
	var result int
	for i := 0; i < len(b); i++ {
		result = result*10 + int(b[i]-'0')
	}
	return result
}
