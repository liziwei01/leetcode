/*
 * @Author: liziwei01
 * @Date: 2022-07-21 18:48:00
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-03 02:51:10
 * @Description: file content
 */
package findSubstring

import (
	"fmt"
	"sort"
	"time"
)

func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	words = checkSameWords(words)

	for i := 0; i != len(s); i++ {
		if judgeSubstring(s[i:], words) {
			ans = append(ans, i)
		}
	}

	return ans
}

// 检查s中是否恰好可以由 words 中所有单词串联形成的子串
// 检查s开头是否正好是某个单词
// 如果是，则检查s的后面是否能串联 words 中所有单词
func judgeSubstring(s string, words []string) bool {
	if len(words) == 0 {
		return true
	}

	for wordIdx, word := range words {
		if len(s) < len(word) {
			continue
		}
		if s[:len(word)] == word {
			newWords := remove(words, wordIdx)
			return judgeSubstring(s[len(word):], newWords)
		}
	}

	return false
}

func findSubstrings(s string, words []string) []int {
	start := time.Now()
	ans := make([]int, 0)
	possibleSubString := permute(words)
	permuteTime := time.Since(start)
	fmt.Println(permuteTime)
	for _, word := range possibleSubString {
		if idxs := newStrStr(s, word); len(idxs) != 0 {
			ans = append(ans, idxs...)
		}
	}
	strTime := time.Since(start)
	fmt.Println(strTime)
	ans = removeDuplicates(ans)
	removeTime := time.Since(start)
	fmt.Println(removeTime)
	return ans
}

func strStr(haystack string, needle string) int {
	needle_len := len(needle)
	if needle_len == 0 {
		return 0
	}

	haystack_len := len(haystack)
	if haystack_len == 0 {
		return -1
	}

	for i := 0; i < haystack_len-needle_len+1; i++ {
		if match(haystack[i:i+needle_len], needle) {
			return i
		}
	}

	return -1
}

func newStrStr(haystack string, needle string) []int {
	ans := make([]int, 0)
	needle_len := len(needle)
	if needle_len == 0 {
		return ans
	}

	haystack_len := len(haystack)
	if haystack_len == 0 {
		return ans
	}

	for i := 0; i < haystack_len-needle_len+1; i++ {
		if match(haystack[i:i+needle_len], needle) {
			ans = append(ans, i)
		}
	}

	return ans
}

func match(str1, str2 string) bool {
	return str1 == str2
}

func permute(nums []string) []string {
	mp := make(map[int]bool)
	tmp := ""
	ans := make([]string, 0)
	var dfs func(now int)
	dfs = func(now int) {
		if now == len(nums) {
			ans = append(ans, tmp)
			return
		}
		for i, j := range nums {
			if mp[i] != true {
				mp[i] = true
				tmp += j
				dfs(now + 1)
				tmp = tmp[:len(tmp)-len(j)]
				mp[i] = false
			}
		}
	}
	dfs(0)
	return ans
}

func remove(words []string, i int) []string {
	newWords := make([]string, len(words))
	copy(newWords, words)
	newWords = append(newWords[:i], newWords[i+1:]...)
	return newWords
}

func removeDuplicates(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return nums
}

func removeDuplicatesStr(nums []string) []string {
	sort.Strings(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return nums
}

// 检查如果所有的单词都相同（比如全是a），则只返回一个单词
func checkSameWords(words []string) []string {
	if len(words) == 0 || len(words) == 1 {
		return words
	}

	example := words[0]
	res := ""
	same := true

	for _, word := range words {
		res += word
		if word != example {
			same = false
			break
		}
	}

	if same {
		return []string{res}
	}

	return words
}
