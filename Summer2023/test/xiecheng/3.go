/*
 * @Author: liziwei01
 * @Date: 2023-10-10 08:11:41
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-10 08:56:24
 * @Description: file content
 */
package xiecheng

import (
	"fmt"
)

// s中的字母可以修改为字母表中相邻的字母如a改为b，b改为a，使s中的相邻字母不相同，输出最少修改次数与一个可能的最终结果
func main() {
	s := ""
	fmt.Scan(&s)
	if len(s) <= 1 {
		fmt.Println(0)
		fmt.Println(s)
		return
	}

	modified := 0
	for i := 1; i != len(s); i++ {
		if s[i-1] == s[i] {
			// 需要处理一次
			modified++
			if s[i] == 'a' {
				// 边界条件a
				if i+1 < len(s) && s[i+1] == 'b' {
					// 路被走死，只能改前一个了
					if i-2 >= 0 {
						// i-1前面还有
						b := differentByte(s[i-1], s[i-2], 1)
						s = s[:i-1] + string(b) + s[i:]
					} else {
						b := differentByte(s[i-1], s[i-1], 1)
						s = s[:i-1] + string(b) + s[i:]
					}
					// 处理完了，不继续往下走
					continue
				} else {
					// 路还有，继续往下面处理走
				}
			} else if s[i] == 'z' {
				// 边界条件z
				if i+1 < len(s) && s[i+1] == 'y' {
					// 路被走死，只能改前一个了
					if i-2 >= 0 {
						// i-1前面还有
						b := differentByte(s[i-1], s[i-2], 1)
						s = s[:i-1] + string(b) + s[i:]
					} else {
						b := differentByte(s[i-1], s[i-1], 1)
						s = s[:i-1] + string(b) + s[i:]
					}
					// 处理完了，不继续往下走
					continue
				} else {
					// 路还有，继续往下面处理走
				}
			}
			if i+1 < len(s) {
				// 后面还有
				b := differentByte(s[i], s[i+1], 1)
				s = s[:i] + string(b) + s[i+1:]
			} else {
				// 后面没了
				b := differentByte(s[i], s[i], 1)
				s = s[:i] + string(b) + s[i+1:]
			}
		}
	}
	fmt.Println(modified)
	fmt.Println(s)
}

// a修改为相邻字母，不能与b相同
func differentByte(a, b, diff byte) byte {
	ret := (a+diff-'a')%26 + 'a'
	if ret == b {
		return differentByte(a, b, -diff)
	}
	return ret
}
