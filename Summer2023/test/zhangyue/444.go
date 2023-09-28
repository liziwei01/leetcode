/*
 * @Author: liziwei01
 * @Date: 2023-09-27 06:50:37
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-27 06:50:37
 * @Description: file content
 */
package zhangyue

func isPalindrome(s string) bool {
	// 检查字符串是否是回文数
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func minPalindromes(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// 创建一个dp数组，dp[i]表示s[0:i]的最小回文分割次数
	dp := make([]int, n)
	for i := range dp {
		dp[i] = i + 1
	}

	for i := 1; i < n; i++ {
		if isPalindrome(s[:i+1]) {
			dp[i] = 1
			continue
		}
		for j := 0; j < i; j++ {
			if isPalindrome(s[j+1 : i+1]) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
