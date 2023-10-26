/*
 * @Author: liziwei01
 * @Date: 2023-10-26 07:17:51
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-26 07:57:58
 * @Description: file content
 */
package shopee

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 *
 * @param strs string字符串 一维数组
 * @return int整型
 */
func numSimilarGroups(strs []string) int {
	sum := 0
	visited := make(map[int]bool)
	for i := 0; i < len(strs); i++ {
		if visited[i] {
			continue
		}
		sum++
		visited[i] = true
		similarGroup := make([]string, 0)
		similarGroup = append(similarGroup, strs[i])
		previousLen := 0
		// 一直扫描直到没有相似的str
		for {
			previousLen = len(similarGroup)
			// 删去所有剩余的strs中的相似的str,从第一个str开始
			for j := i + 1; j < len(strs); j++ {
				if !visited[j] && similar(similarGroup, strs[j]) {
					visited[j] = true
					similarGroup = append(similarGroup, strs[j])
				}
			}
			if previousLen == len(similarGroup) {
				break
			}
		}
	}
	return sum
}

// 任何一个str和str2相似都return true
func similar(strs []string, str2 string) bool {
	for _, str := range strs {
		if isSimilar(str, str2) {
			return true
		}
	}
	return false
}

// 判断str和str2是否相似
func isSimilar(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}
	count := 0
	for i := 0; i != len(str1); i++ {
		if str1[i] != str2[i] {
			count++
		}
	}

	return count == 2
}
