/*
 * @Author: liziwei01
 * @Date: 2022-06-16 07:49:16
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-16 07:56:54
 * @Description: file content
 */
package romanToInt

func romanToInt(s string) int {
	length := len(s)
	ret := 0

	for i := 0; i < length; i++ {
		if s[i] == 'I' && i+1 < length && s[i+1] == 'V' {
			ret += 4
			i++
		} else if s[i] == 'I' && i+1 < length && s[i+1] == 'X' {
			ret += 9
			i++
		} else if s[i] == 'X' && i+1 < length && s[i+1] == 'L' {
			ret += 40
			i++
		} else if s[i] == 'X' && i+1 < length && s[i+1] == 'C' {
			ret += 90
			i++
		} else if s[i] == 'C' && i+1 < length && s[i+1] == 'D' {
			ret += 400
			i++
		} else if s[i] == 'C' && i+1 < length && s[i+1] == 'M' {
			ret += 900
			i++
		} else if s[i] == 'I' {
			ret += 1
		} else if s[i] == 'V' {
			ret += 5
		} else if s[i] == 'X' {
			ret += 10
		} else if s[i] == 'L' {
			ret += 50
		} else if s[i] == 'C' {
			ret += 100
		} else if s[i] == 'D' {
			ret += 500
		} else if s[i] == 'M' {
			ret += 1000
		}
	}

	return ret
}
