/*
 * @Author: liziwei01
 * @Date: 2022-08-03 03:15:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-12 16:59:45
 * @Description: file content
 */
package longestValidParentheses

func longestValidParentheses(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return 0
	}

	// 先去除所有开头的右括号
	for i := 0; i < sLen; i++ {
		if s[0] == ')' {
			s = s[1:]
			sLen--
		}
	}

	if sLen == 0 {
		return 0
	}

	// 再去除所有结尾的左括号
	for i := sLen - 1; i > -1; i-- {
		if s[sLen-1] == '(' {
			s = s[:sLen-1]
			sLen--
		}
	}

	if sLen == 0 {
		return 0
	}

	// 开始计算最长的有效括号
	// 1 如果能直接匹配成功最好
	var h = heap{
		data: make([]dataStruct, 0),
		size: 0,
	}

	for i := 0; i != sLen; i++ {
		if s[i] == '(' {
			h.push(
				dataStruct{
					r: '(',
					i: i,
				},
			)
		} else {
			if d := h.pop(); d.r != '(' {
				// 2 如果发现左括号用完了，说明右括号多了。前半【(())  )  ()】是有效的，再来单独看看后半
				return max(i, longestValidParentheses(s[i+1:]))
			}
		}
	}

	if h.size == 0 {
		return sLen
	}

	// 3 如果是左括号多了，那么就开始分段，每个多余的左括号【()  (  ()()  (  (())】所在的位置都是一个楔子，把str分成很多段重新看
	maximum := longestValidParentheses(s[:h.data[0].i])
	for idx := 0; idx < h.size-1; idx++ {
		left := h.data[idx].i
		right := h.data[idx+1].i
		maximum = max(maximum, longestValidParentheses(s[left+1:right]))
	}

	maximum = max(maximum, longestValidParentheses(s[h.data[h.size-1].i+1:]))

	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type dataStruct struct {
	r rune
	i int
}

type heap struct {
	data []dataStruct
	size int
}

func (h *heap) push(r dataStruct) {
	h.data = append(h.data, r)
	h.size++
}

func (h *heap) pop() dataStruct {
	if h.size == 0 {
		return dataStruct{
			r: 0,
			i: -1,
		}
	}

	h.size--
	r := h.data[h.size]
	h.data = h.data[:h.size]
	return r
}

func (h *heap) top() dataStruct {
	if h.size == 0 {
		return dataStruct{
			r: 0,
			i: -1,
		}
	}

	return h.data[h.size-1]
}

func (h *heap) bottom() dataStruct {
	if h.size == 0 {
		return dataStruct{
			r: 0,
			i: -1,
		}
	}

	return h.data[0]
}
