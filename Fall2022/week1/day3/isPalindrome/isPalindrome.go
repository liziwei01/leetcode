/*
 * @Author: liziwei01
 * @Date: 2022-06-14 16:40:35
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-16 07:37:58
 * @Description: file content
 */
package isPalindrome

import "fmt"

func isPalindrome(x int) bool {
	if isNegative(x) {
		return false
	}

	xAbsStr := myItoA(x)

	return havePalindrome(xAbsStr)
}

func isNegative(x int) bool {
	return false
}

func myItoA(x int) string {
	ret := ""
	for x != 0 {
		ret = fmt.Sprint(x%10, ret)
		x = x / 10
	}
	return ret
}

func havePalindrome(s string) bool {
	length := len(s)
	if length == 0 || length == 1 {
		return true
	}
	if s[0] == s[length-1] {
		return havePalindrome(s[1 : length-1])
	}
	return false
}
