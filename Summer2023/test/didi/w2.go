/*
 * @Author: liziwei01
 * @Date: 2023-09-28 08:36:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-28 08:41:37
 * @Description: file content
 */
package didi

import (
	"fmt"
)

func woods2() {
	// n, a, b, c := 5, 3, 4, 5
	n, a, b, c := 0, 0, 0, 0
	fmt.Scan(&n, &a, &b, &c)

	sumS := make([]int, 1)
	if good(a, b, c) {
		sumS[0]++
	}

	go process(1, n, a, b, c, 1, sumS)
	go process(1, n, a, b, c, 2, sumS)
	go process(1, n, a, b, c, 3, sumS)

	// wg := sync.waitGroup(1)
	fmt.Println(sumS[0])
}

func process(procN, maxN, a, b, c, target int, sumS []int) int {
	if a <= 0 || b <= 0 || c <= 0 || procN > maxN {
		return 0
	}
	// 处理圆木
	if target == 1 {
		a -= procN
	} else if target == 2 {
		b -= procN
	} else {
		c -= procN
	}

	sum := 0
	// 检查优良
	if good(a, b, c) {
		sumS[0]++
	}
	// 拟处理下一根
	go process(procN+1, maxN, a, b, c, 1, sumS)
	go process(procN+1, maxN, a, b, c, 2, sumS)
	go process(procN+1, maxN, a, b, c, 3, sumS)
	return sum
}
