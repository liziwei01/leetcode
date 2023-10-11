/*
 * @Author: liziwei01
 * @Date: 2023-10-10 07:02:52
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-10 08:42:26
 * @Description: file content
 */
package xiecheng

import (
	"fmt"
)

// s修改一个字母，使之包含尽可能多的长度为3的回文子串
func pa() {
	s := "01101"
	// fmt.Scan(&s)

	sum := 0
	for i := 0; i != len(s); i++ {
		tmpS := palinNum(s[:i] + "0" + s[i+1:])
		if tmpS > sum {
			sum = tmpS
		}
		tmpS = palinNum(s[:i] + "1" + s[i+1:])
		if tmpS > sum {
			sum = tmpS
		}
	}

	fmt.Println(sum)
}

func palinNum(s string) int {
	length := len(s)
	sum := 0
	i := 0
	for i < length-2 {
		switch s[i : i+3] {
		case "000", "010", "101", "111":
			sum++
			i++
		default:
			i++
		}
	}
	return sum
}

// 011011 011111
// 00000
// 00001
// 00010 .
// 00011
// 00100
// 00101
// 00110
// 00111
// 01000
// 01001
// 01010
// 01011
// 01100
// 01101
// 01110
// 01111
// 10000
// 10001
// 10010
// 10011
// 10100
// 10101
// 10110
// 10111
// 11000
// 11001
// 11010
// 11011
// 11100
// 11101
// 11110
// 11111
