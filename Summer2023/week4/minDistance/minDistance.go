/*
 * @Author: liziwei01
 * @Date: 2023-07-10 10:36:15
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-11 00:43:19
 * @Description: file content
 */
package mindistance

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	m := len(word1)
	n := len(word2)
	if m == 0 || n == 0 {
		return max(m, n)
	}
	distances := make([][]int, 0)
	// 每个格子代表着单词的前i个字母和另一个单词前j个字母的编辑距离
	// initialize
	for i := 0; i != m+1; i++ {
		distance := make([]int, n+1)
		distance[0] = i
		distances = append(distances, distance)
	}
	for i := 0; i != n+1; i++ {
		distances[0][i] = i
	}

	for i := 1; i != m+1; i++ {
		for j := 1; j != n+1; j++ {
			distances[i][j] = min(distances[i-1][j]+1, distances[i][j-1]+1, distances[i-1][j-1]+change(word1, word2, i-1, j-1))
		}
	}

	return distances[m][n]
}

// 当前字母是否需要修改以使当前字母相等
func change(word1 string, word2 string, i, j int) int {
	if word1[i] == word2[j] {
		return 0
	}

	return 1
}

func max(m, n int) int {
	if m < n {
		return n
	}
	return m
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	} else {
		return c
	}
}
