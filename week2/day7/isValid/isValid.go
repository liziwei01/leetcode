/*
 * @Author: liziwei01
 * @Date: 2022-06-29 01:39:42
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-29 01:51:34
 * @Description: file content
 */
package isValid

type heap struct {
	data []rune
	size int
}

func isValid(s string) bool {
	length := len(s)

	if length == 0 {
		return true
	}

	if length%2 != 0 {
		return false
	}

	h := heap{
		data: make([]rune, 0),
		size: 0,
	}

	for i := 0; i != length; i++ {
		if s[i] == '(' {
			h.push('(')
		} else if s[i] == '[' {
			h.push('[')
		} else if s[i] == '{' {
			h.push('{')
		} else if s[i] == ')' {
			if h.pop() != '(' {
				return false
			}
		} else if s[i] == ']' {
			if h.pop() != '[' {
				return false
			}
		} else if s[i] == '}' {
			if h.pop() != '{' {
				return false
			}
		} else {
			return false
		}
	}

	if h.size == 0 {
		return true
	}

	return false
}

func (h *heap) push(r rune) {
	h.data = append(h.data, r)
	h.size++
}

func (h *heap) pop() rune {
	if h.size == 0 {
		return 0
	}

	h.size--
	r := h.data[h.size]
	h.data = h.data[:h.size]
	return r
}
