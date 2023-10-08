/*
 * @Author: liziwei01
 * @Date: 2023-10-08 07:37:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-08 08:12:12
 * @Description: file content
 */
package webank

import (
	"fmt"
	"strings"
)

// 塔吊与旧楼距离需要超过R，与新楼需要小于R，几个地点可以建塔吊？TLE
func buildingsq() {
	n, m, R := 0, 0, 0
	fmt.Scan(&n, &m, &R)
	space := make([]string, n)
	x, y := 0, 0
	for i := 0; i != n; i++ {
		fmt.Scan(&space[i])
		index := strings.Index(space[i], "2")
		if index >= 0 {
			x, y = i, index
		}
	}

	num := 0
	x0 := x - R
	if x0 < 0 {
		x0 = 0
	}
	x1 := x + R
	if x1 > n - 1 {
		x1 = n - 1
	}
	y0 := y - R
	if y0 < 0 {
		y0 = 0
	}
	y1 := y + R
	if y1 > m - 1 {
		y1 = m - 1
	}
	for i := x0; i <= x1; i++ {
		for j := y0; j != y1; j++ {
			if canReachNew(space, i, j, x, y, R) && cannotReachOld(space, i, j, n, m, R) {
				num++
			}
		}
	}
	fmt.Println(num)
}

func canReachNew(space []string, i, j, x, y, R int) bool {
	return abs(i-x) + abs(j-y) <= R
}

func cannotReachOld(space []string, i, j, n, m, R int) bool {
	x0 := i - R
	if x0 < 0 {
		x0 = 0
	}
	x1 := i + R
	if x1 > n - 1 {
		x1 = n - 1
	}
	y0 := j - R
	if y0 < 0 {
		y0 = 0
	}
	y1 := j + R
	if y1 > m - 1 {
		y1 = m - 1
	}
	for a := x0; a <= x1; a++ {
		for b := y0; b <= y1; b++ {
			if abs(a-i) + abs(b-j) <= R && space[a][b] == '1' {
				return false
			}
		}
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}