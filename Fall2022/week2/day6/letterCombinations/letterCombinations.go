/*
 * @Author: liziwei01
 * @Date: 2022-06-28 14:11:19
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-28 15:08:49
 * @Description: file content
 */
package letterCombinations

import "sort"

var (
	m = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var tmp = &[]string{}
	possibleLetters := m[digits[0]]
	possibleLettersLength := len(possibleLetters)
	for i := 0; i != possibleLettersLength; i++ {
		*tmp = append(*tmp, possibleLetters[i])
	}

	letterComb(digits[1:], tmp)

	return *tmp
}

func letterComb(digits string, tmp *[]string) {
	if len(digits) == 0 {
		return
	}

	originalRes := deepCopy(tmp)
	originalResLength := len(originalRes)
	possibleLetters := m[digits[0]]
	possibleLettersLength := len(possibleLetters)
	for i := 0; i != possibleLettersLength-1; i++ {
		*tmp = append(*tmp, originalRes...)
	}

	sort.Strings(*tmp)

	for i := 0; i != originalResLength; i++ {
		for j := 0; j != possibleLettersLength; j++ {
			(*tmp)[i*possibleLettersLength+j] += possibleLetters[j]
		}
	}

	letterComb(digits[1:], tmp)
}

func deepCopy(tmp *[]string) []string {
	res := make([]string, len(*tmp))
	copy(res, *tmp)
	return res
}
