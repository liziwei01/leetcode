/*
 * @Author: liziwei01
 * @Date: 2022-06-14 12:07:50
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-14 14:23:06
 * @Description: file content
 */
package longestpalindrome

func longestPalindrome(s string) string {
	length := len(s)
	result := s[0:1]
	if length == 0 || length == 1 {
		return s
	}
	for i := 0; i < length; i++ {
		if len(result) >= length-i {
			break
		}
		for j := length; j > i+1; j-- {
			if len(result) >= j-i {
				break
			}
			if havePalindrome(s[i:j]) {
				if len(s[i:j]) > len(result) {
					result = s[i:j]
				}
			}
		}
	}
	return result
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
