/*
 * @Author: liziwei01
 * @Date: 2023-06-09 00:02:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-09 19:07:01
 * @Description: file content
 */
package ismatch

// s string, p pattern
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
		if p[0] == '*' {
			return isMatch(s, p[1:])
		}
		return false
	}

	// 如果s没走完，p走完了
	if len_p == 0 {
		return false
	}

	// 如果p的第一个字符是*
	if p[0] == '*' {
		// 如果下一个还是*就合并
		if len_p > 1 && p[1] == '*' {
			return isMatch(s, p[1:])
		}
		return isMatch(s, p[1:]) || isMatch(s[1:], p) || isMatch(s[1:], p[1:])
	}

	// 如果p的第一个字符不是*
	if p[0] == '?' || p[0] == s[0] {
		return isMatch(s[1:], p[1:])
	}

	return false
}
