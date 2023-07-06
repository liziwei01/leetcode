/*
 * @Author: liziwei01
 * @Date: 2023-06-09 00:02:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-10 22:26:36
 * @Description: file content
 */
package ismatch

// expression
type expr struct {
	pattern string
}

func isMatch(s string, p string) bool {
	// 排除p为空
	if len(p) == 0 {
		return s == p
	}
	// 编译 pattern
	longExpr, hasStar, hasPrefixStar, hasSuffixStar := compile(p)
	if longExpr.equal(s) {
		return true
	}
	exprs := longExpr.simplify()

	// 排除s为空
	if len(exprs) == 0 {
		return hasStar
	} else if len(s) == 0 && len(exprs) != 0 {
		return false
	}

	// 仅有一个pattern
	if len(exprs) == 1 {
		fIdx := exprs[0].matchFirst(s)
		if fIdx == -1 {
			return false
		}
		lIdx := exprs[0].matchLast(s)
		if hasPrefixStar && hasSuffixStar {
			// 1 *?* 两边都有*
			return true
		} else if hasPrefixStar && !hasSuffixStar {
			// 2 *?
			return (lIdx + exprs[0].len()) == len(s)
		} else if !hasPrefixStar && hasSuffixStar {
			// 3 ?*
			return fIdx == 0
		} else {
			// 4 ?
			return false
		}
	}

	fIdxNext := exprs[0].matchFirstforNextIdx(s)
	if fIdxNext == -1 {
		return false
	}

	lIdxNext := exprs[len(exprs)-1].matchLastforNextIdx(s)
	if lIdxNext == -1 {
		return false
	}

	newS := ""
	if fIdxNext > lIdxNext {
		return false
	}
	if fIdxNext != lIdxNext {
		newS = s[fIdxNext:lIdxNext]
	}

	newExprs := make([]*expr, 0)
	if 1 < len(exprs)-1 {
		newExprs = exprs[1 : len(exprs)-1]
	}

	b := ismatch(newS, newExprs, hasStar)
	if !b {
		return false
	}

	return isFullMatch(len(s), fIdxNext-exprs[0].len(), lIdxNext+exprs[len(exprs)-1].len(), hasPrefixStar, hasSuffixStar)
}

func isFullMatch(slen, strMatchStart, strMatchEnd int, hasPrefixStar, hasSuffixStar bool) bool {
	b := true
	if strMatchStart != 0 {
		b = b && hasPrefixStar
	}
	if strMatchEnd != slen {
		b = b && hasSuffixStar
	}
	return b
}

func ismatch(s string, exprs []*expr, hasStar bool) bool {
	if len(exprs) == 0 {
		return hasStar
	} else if len(s) == 0 && len(exprs) != 0 {
		return false
	}

	if len(exprs) == 1 {
		if len(s) == exprs[0].len() {
			return exprs[0].matchFirst(s) != -1
		}
		return exprs[0].matchFirst(s) != -1 && hasStar
	}

	fIdxNext := exprs[0].matchFirstforNextIdx(s)
	if fIdxNext == -1 {
		return false
	}

	lIdxNext := exprs[len(exprs)-1].matchLastforNextIdx(s)
	if lIdxNext == -1 {
		return false
	}

	newS := ""
	if fIdxNext > lIdxNext {
		return false
	}
	if fIdxNext != lIdxNext {
		newS = s[fIdxNext:lIdxNext]
	}

	newExprs := make([]*expr, 0)
	if 1 < len(exprs)-1 {
		newExprs = exprs[1 : len(exprs)-1]
	}

	return ismatch(newS, newExprs, hasStar)
}

// 去除多余的*
func compile(p string) (*expr, bool, bool, bool) {
	hasStar := false
	isStar := false
	pat := ""
	for i := 0; i != len(p); i++ {
		if p[i] != '*' {
			// 如果是正常英文字母或?，保留该字母
			pat += string(p[i])
			isStar = false
		} else if p[i] == '*' && !isStar {
			// 如果是*且前一位不是*，保留且记录该位是*
			pat += string(p[i])
			hasStar = true
			isStar = true
		} else {
			// 如果是*且前一位是*，不保留
		}
	}

	return &expr{
		pattern: pat,
	}, hasStar, p[0] == '*', p[len(p)-1] == '*'
}

// 去除所有的*
func (e *expr) simplify() []*expr {
	exprs := make([]*expr, 0)
	pat := ""

	for i := 0; i != len(e.pattern); i++ {
		if e.pattern[i] != '*' {
			pat += string(e.pattern[i])
		} else {
			if i != 0 {
				exprs = append(exprs, &expr{pattern: pat})
				pat = ""
			}
		}
	}
	if len(pat) != 0 {
		exprs = append(exprs, &expr{pattern: pat})
	}

	return exprs
}

func (e *expr) canMatch(s string) bool {
	return len(e.pattern) <= len(s)
}

func (e *expr) len() int {
	return len(e.pattern)
}

func (e *expr) matchFirstforNextIdx(s string) int {
	idx := e.matchFirst(s)
	if idx == -1 {
		return -1
	}

	return idx + e.len()
}

func (e *expr) matchLastforNextIdx(s string) int {
	return e.matchLast(s)
}

func (e *expr) matchFirst(s string) int {
	if !e.canMatch(s) {
		return -1
	}

	for i := 0; i != len(s)-e.len()+1; i++ {
		if e.equal(s[i : i+e.len()]) {
			return i
		}
	}

	return -1
}

func (e *expr) matchLast(s string) int {
	if !e.canMatch(s) {
		return -1
	}

	idxes := e.match(s)

	if len(idxes) == 0 {
		return -1
	}

	return idxes[len(idxes)-1]
}

// return 所有match的substring的idx
func (e *expr) match(s string) []int {
	idxes := make([]int, 0)

	for i := 0; i != len(s)-e.len()+1; i++ {
		if e.equal(s[i : i+e.len()]) {
			idxes = append(idxes, i)
		}
	}

	return idxes
}

func (e *expr) equal(s string) bool {
	if len(s) != e.len() {
		return false
	}

	if s == e.pattern {
		return true
	}

	for i := 0; i != len(s); i++ {
		if s[i] != e.pattern[i] && e.pattern[i] != '?' {
			return false
		}
	}

	return true
}