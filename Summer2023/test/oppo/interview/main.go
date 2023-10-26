/*
 * @Author: liziwei01
 * @Date: 2023-10-26 00:35:40
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-26 00:35:41
 * @Description: file content
 */
package main

import "fmt"

/**
 * @description: main
 * @param {*}
 * @return {*}
 */
func main() {
	fmt.Println(maxSubStr("abcdabcde"))
}

func maxSubStr(str string) int {
	length := len(str)
	m := make(map[byte]int)
	start := 0
	end := 0
	maxLength := 0
	tmpMaxLength := 0
	// 滑动窗口 abcaabcde
	for start < length && end < length {
		b := str[end]
		if m[b] == 0 {
			// 不存在就扩大窗口
			m[b] = end
			end++
			tmpMaxLength++
		} else {
			// 存在
			// 当前窗口大小是否是所有扫到的窗口中最大的
			if tmpMaxLength > maxLength {
				maxLength = tmpMaxLength
			}
			// 窗口从重复的字符的+1处开始新建
			start = m[b] + 1
			end = m[b] + 1
			m = make(map[byte]int)
			tmpMaxLength = 0
		}
	}
	return maxLength
}
