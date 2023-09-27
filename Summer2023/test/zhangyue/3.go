/*
 * @Author: liziwei01
 * @Date: 2023-09-27 06:37:05
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-27 06:38:15
 * @Description: file content
 */

package zhangyue

import (
	"fmt"
	"strings"
)

func m3() {
	s := ""
	fmt.Scan(&s)

	num1 := try(s)

	fmt.Print(num1)
}

func try(s string) int {
	num := -1
	i := 0
	tmpS := s
	for i < len(s) {
		if isSyn(tmpS) {
			i += len(tmpS)
			fmt.Println(tmpS)
			tmpS = rev(s[i:])
			num++
		} else {
			tmpS = tmpS[:len(tmpS)-1]
		}
	}
	return num
}

func isSyn(s string) bool {
	for i := 0; i != len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func rev(s string) string {
	newS := make([]string, len(s))
	for i := 0; i != len(s)/2; i++ {
		newS[i], newS[len(newS)-1-i] = string(s[len(newS)-1-i]), string(s[i])
	}
	return strings.Join(newS, "")
}
