/*
 * @Author: liziwei01
 * @Date: 2023-07-08 02:40:39
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 02:41:46
 * @Description: file content
 */
package lengthoflastword

func lengthOfLastWord(s string) int {
	wordLen := 0
	findWord := false
	for i := len(s) - 1; i != 0; i-- {
		if s[i] == ' ' && !findWord {
			continue
		} else if s[i] == ' ' && findWord {
			break
		} else {
			wordLen++
			findWord = true
		}
	}

	return wordLen
}