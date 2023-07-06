/*
 * @Author: liziwei01
 * @Date: 2023-06-10 10:17:48
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-10 10:17:49
 * @Description: file content
 */
package ismatch

func ismatchs(s string, p string) bool {
	return match(s, p, noStarLen(p))
}

// s string, p pattern
func match(s string, p string, nlen int) bool {
	len_s := len(s)
	len_p := len(p)

	if nlen > len_s {
		return false
	}

	// 如果s走完了
	if len_s == 0 {
		// 如果p也走完了
		if len_p == 0 {
			return true
		}
		// 如果p还没走完
		if p[0] == '*' {
			return match(s, p[1:], nlen-1)
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
			return match(s, p[1:], nlen-1)
		}
		return match(s, p[1:], nlen-1) || match(s[1:], p, nlen) || match(s[1:], p[1:], nlen-1)
	}

	// 如果p的第一个字符不是*
	if p[0] == '?' || p[0] == s[0] {
		return match(s[1:], p[1:], nlen)
	}

	return false
}

func noStarLen(s string) int {
	n := 0
	for i := 0; i != len(s); i++ {
		if s[i] != '*' {
			n++
		}
	}
	return n
}
