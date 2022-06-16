/*
 * @Author: liziwei01
 * @Date: 2022-06-15 15:41:17
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-16 07:38:04
 * @Description: file content
 */
package isMatch

func isMatch(s string, p string) bool {
	len_s := len(s)
	len_p := len(p)

	// 如果s走完了
	if len_s == 0 {
		// 如果p也走完了
		if len_p == 0 {
			return true
		}
		// 如果p还没走完
		if len_p > 1 && p[1] == '*' {
			return isMatch(s, p[2:])
		}
		return false
	}

	// 如果p走完了
	if len_p == 0 {
		return false
	}

	// 如果p的第二个字符是*
	if len_p > 1 && p[1] == '*' {
		// 如果p的第一个字符和s的第一个字符匹配
		if p[0] == '.' || p[0] == s[0] {
			return isMatch(s, p[2:]) || isMatch(s[1:], p) || isMatch(s[1:], p[2:])
		}
		// 如果p的第一个字符和s的第一个字符不匹配，说明*可以匹配0个字符
		return isMatch(s, p[2:])
	}

	// 如果p的第二个字符不是*
	if p[0] == '.' || p[0] == s[0] {
		return isMatch(s[1:], p[1:])
	}

	return false
}
