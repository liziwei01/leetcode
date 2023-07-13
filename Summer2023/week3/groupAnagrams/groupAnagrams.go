/*
 * @Author: liziwei01
 * @Date: 2023-07-04 17:44:32
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-11 21:22:08
 * @Description: file content
 */
package groupanagrams

import (
	"sort"
)

type stringIdx struct {
	str string
	idx int
}

func groupAnagrams(strs []string) [][]string {
	mid := deepCopyAddIdx(strs)
	intraOrder(mid)
	interOrder(mid)

	ret := make([][]string, 0)
	temp := mid[0]
	ret = append(ret, []string{strs[temp.idx]})
	retIdx := 0
	for i := 1; i < len(mid); i++ {
		if mid[i].str == temp.str {
			ret[retIdx] = append(ret[retIdx], strs[mid[i].idx])
		} else {
			temp = mid[i]
			ret = append(ret, []string{strs[temp.idx]})
			retIdx++
		}
	}

	return ret
}

func deepCopyAddIdx(strs []string) []stringIdx {
	ret := make([]stringIdx, len(strs))
	for i := 0; i != len(strs); i++ {
		ret[i] = stringIdx{strs[i], i}
	}
	return ret
}

func intraOrder(strs []stringIdx) {
	for i := 0; i != len(strs); i++ {
		chars := []rune(strs[i].str)
		sort.Slice(chars, func(i int, j int) bool {
			return chars[i] < chars[j]
		})
		strs[i].str = string(chars)
	}
}

func interOrder(strs []stringIdx) {
	sort.Slice(strs, func(i int, j int) bool {
		return strs[i].str < strs[j].str
	})
}
