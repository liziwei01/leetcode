/*
 * @Author: liziwei01
 * @Date: 2023-09-15 09:07:41
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-15 09:24:55
 * @Description: file content
 */
package tencent2

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	T, n := 0, 0
	fmt.Scan(&T)
	for i := 0; i != T; i++ {
		fmt.Scan(&n)
		stringSlice := make([]string, n)
		for j := 0; j != n; j++ {
			fmt.Scan(&stringSlice[j])
		}
		if hasSpine(stringSlice) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func hasSpine(strs []string) bool {
	n := len(strs)
	for i := 0; i != n; i++ {
		tmp := strings.Split(strs[i], "")
		sort.Strings(tmp)
		strs[i] = strings.Join(tmp, "")
	}
	sort.Strings(strs)
	for i := 0; i < n-1; i++ {
		if strs[i] == strs[i+1] {
			return true
		}
	}
	return false
}
