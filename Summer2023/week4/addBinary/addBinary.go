/*
 * @Author: liziwei01
 * @Date: 2023-07-09 00:41:43
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-09 00:53:24
 * @Description: file content
 */
package addbinary

import "fmt"

func addBinary(a string, b string) string {
	al := len(a)
	bl := len(b)
	s := ""
	overflow := 0

	if al < bl {
		a = zeros(bl-al) + a
	} else {
		b = zeros(al-bl) + b
	}

	for i := len(a) - 1; i >= 0; i-- {
		sum := toD(a[i]) + toD(b[i]) + overflow
		s = fmt.Sprint(sum%2) + s
		overflow = sum / 2
	}
	if overflow != 0 {
		s = fmt.Sprint(overflow) + s
	}

	return s
}

func zeros(n int) string {
	s := ""
	for i := 0; i != n; i++ {
		s += "0"
	}
	return s
}

func toD(b byte) int {
	return int(b - '0')
}
