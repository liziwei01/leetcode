/*
 * @Author: liziwei01
 * @Date: 2023-07-08 23:35:21
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-09 00:46:45
 * @Description: file content
 */
package isnumber

import "strings"

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 全部小写，去除符号
	s = strings.ToLower(s)
	if bIsSign(s[0]) {
		return isNum(s[1:])
	}
	return isNum(s)
}

// 无符号小数或整数后面（可选）跟e与整数
func isNum(s string) bool {
	if len(s) == 0 {
		return false
	}
	isDecimal := false
	for i := 0; i != len(s); i++ {
		// 整数
		if bIsDigit(s[i]) {
			continue
		} else if bIsE(s[i]) {
			// e后只能是无符号整数
			if i == 0 {
				return false
			}
			return IsSignedInt(s[i+1:])
		} else if bIsDot(s[i]) {
			if isDecimal {
				return false
			}
			isDecimal = true
			if i == 0 {
				// 一个点 '.' ，后面跟着至少一位数字
				if i == len(s)-1 || !IsFullInt(s[1:2]) {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}

// 有符号整数
func IsSignedInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	if bIsSign(s[0]) {
		s = s[1:]
	}
	return IsFullInt(s)
}

// 不带任何别的东西的纯数字
func IsFullInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i != len(s); i++ {
		if !bIsDigit(s[i]) {
			return false
		}
	}

	return true
}

func bIsSign(b byte) bool {
	return b == '+' || b == '-'
}

func bIsDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}

	return false
}

func bIsE(b byte) bool {
	return b == 'e'
}

func bIsDot(b byte) bool {
	return b == '.'
}

func toD(b byte) int {
	return int(b - '0')
}