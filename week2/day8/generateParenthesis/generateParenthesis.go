/*
 * @Author: liziwei01
 * @Date: 2022-06-30 03:34:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-30 04:06:08
 * @Description: file content
 */
package generateParenthesis

func generateParenthesis(n int) []string {
	// 生成所有可能的排列
	allPossible := gP("", n, n)

	res := make([]string, 0)
	for _, v := range allPossible {
		if isValid(v) {
			res = append(res, v)
		}
	}

	return res
}

func gP(tmp string, leftn, rightn int) []string {
	if leftn == 0 && rightn == 0 {
		return []string{tmp}
	}

	res1 := make([]string, 0)
	if leftn != 0 {
		res1 = gP(tmp+"(", leftn-1, rightn)
	}

	res2 := make([]string, 0)
	if rightn != 0 {
		res2 = gP(tmp+")", leftn, rightn-1)
	}

	res1 = append(res1, res2...)
	return res1
}

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
