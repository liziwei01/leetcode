/*
 * @Author: liziwei01
 * @Date: 2023-09-23 07:07:09
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-23 07:08:51
 * @Description: file content
 */
package oppo

import (
	"fmt"
)

func str() {
	// 两字符串st二选一是否能生成target
	n := 0
	s, t, target := "", "", ""
	fmt.Scan(&n, &s, &t, &target)
	achieved := true
	for i := 0; i != n; i++ {
		if s[i] != target[i] && t[i] != target[i] {
			achieved = false
			break
		}
	}
	if achieved {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
