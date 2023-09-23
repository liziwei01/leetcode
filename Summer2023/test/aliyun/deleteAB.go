/*
 * @Author: liziwei01
 * @Date: 2023-09-17 03:41:45
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-17 03:41:46
 * @Description: file content
 */

package aliyun

import (
	"fmt"
	"strings"
)

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	a := ""
	fmt.Scan(&a)
	newStr := a
	i := strings.Index(newStr, "ab")
	for i != -1 {
		if i < len(newStr)-2 {
			newStr = newStr[:i] + newStr[i+2:]
		} else {
			newStr = newStr[:i]
		}
		i = strings.Index(newStr, "ab")
	}
	limit := 0
	i = strings.Index(newStr, "ba")
	for i != -1 {
		if limit == k {
			break
		}
		if i < len(newStr)-2 {
			newStr = newStr[:i] + newStr[i+2:]
		} else {
			newStr = newStr[:i]
		}
		i = strings.Index(newStr, "ba")
		limit++
	}
	if len(newStr) == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(newStr)
	}
}
