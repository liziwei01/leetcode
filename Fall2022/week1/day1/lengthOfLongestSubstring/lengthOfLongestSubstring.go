/*
 * @Author: liziwei01
 * @Date: 2022-06-13 16:00:16
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-16 07:38:50
 * @Description: file content
 */
package lengthOfLongestSubstring

func lengthOfLongestSubstring1(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	max := 0

	for i := 0; i < len(s); i++ {
		findDup := false
		alphabet := make(map[byte]int)
		for j := i + 1; j < len(s); j++ {
			_, ok := alphabet[s[j]]
			// if find duplicate, break
			if s[i] == s[j] || ok {
				max = Max(max, j-i)
				findDup = true
				break
			}
			alphabet[s[j]] = 1
		}
		// if do not find outer duplicate
		if !findDup {
			max = Max(max, len(s)-i)
			break
		}
	}
	return max
}

func Max(a ...int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
