/*
 * @Author: liziwei01
 * @Date: 2023-09-28 08:09:39
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-28 08:33:45
 * @Description: file content
 */
package didi

import (
	"fmt"
)

func woods() {
	// n, a, b, c := 5, 3, 4, 5
	n, a, b, c := 0, 0, 0, 0
	fmt.Scan(&n, &a, &b, &c)

	sum := 0
	if good(a, b, c) {
		sum++
	}

	var dfs func(procN, target int)
	dfs = func(procN, target int) {
		if a <= 0 || b <= 0 || c <= 0 || procN > n {
			return
		}
		// 处理圆木
		if target == 1 {
			a -= procN
		} else if target == 2 {
			b -= procN
		} else {
			c -= procN
		}

		// 检查优良
		if good(a, b, c) {
			sum++
		}
		// 拟处理下一根
		dfs(procN+1, 1)
		dfs(procN+1, 2)
		dfs(procN+1, 3)

		// 倒退处理
		if target == 1 {
			a += procN
		} else if target == 2 {
			b += procN
		} else {
			c += procN
		}
	}
	go dfs(1, 1)
	go dfs(1, 2)
	go dfs(1, 3)

	fmt.Println(sum)
}

// 能否组成三角形
func good(a, b, c int) bool {
	if a <= b && a <= c {
		if b < c {
			// pass
		} else {
			b, c = c, b
		}
	} else if b <= a && b <= c {
		if a <= c {
			a, b = b, a
		} else {
			a, b, c = b, c, a
		}
	} else {
		if a <= b {
			a, b, c = c, a, b
		} else {
			a, c = c, a
		}
	}
	// 两短边之和大于第三边
	return a+b > c
}
