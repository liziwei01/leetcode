/*
 * @Author: liziwei01
 * @Date: 2022-07-21 18:06:55
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-21 18:15:13
 * @Description: file content
 */
package strStr

func strStr(haystack string, needle string) int {
	needle_len := len(needle)
	if needle_len == 0 {
		return 0
	}

	haystack_len := len(haystack)
	if haystack_len == 0 {
		return -1
	}

	for i := 0; i < haystack_len - needle_len + 1; i++ {
		if match(haystack[i:i+needle_len], needle) {
			return i
		}
	}

	return -1
}

func match(str1, str2 string) bool {
	return str1 == str2
}
