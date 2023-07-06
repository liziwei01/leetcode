/*
 * @Author: liziwei01
 * @Date: 2023-06-10 23:49:07
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-11 00:06:48
 * @Description: file content
 */
package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	le := 0
	m := make(map[byte]int, 0)
	max := 0

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			if max < le {
				max = le
			}
			delete(m, s[i])
			le = 0
		} else {
			m[s[i]] = i
			le++
		}
	}

	if le != 0 {
		if max < le {
			max = le
		}
	}

	return max
}
