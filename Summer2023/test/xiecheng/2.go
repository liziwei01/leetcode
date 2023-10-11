/*
 * @Author: liziwei01
 * @Date: 2023-10-10 07:35:24
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-10 08:08:58
 * @Description: file content
 */
package xiecheng

import (
	"fmt"
	"sort"
)

type point struct {
	x int
	y int
}

// AC
func po() {
	n, x := 0, 0
	fmt.Scan(&n, &x)
	x--

	points := make([]point, n)
	for i := 0; i != n; i++ {
		fmt.Scan(&points[i].x, &points[i].y)
	}
	// points[0] = point{-1, 3}
	// points[1] = point{3, 5}
	// points[2] = point{2, 6}
	// points[3] = point{4, 3}
	fpoint := points[x]

	step := 0
	onLeft := false
	var dfs func(tmpPoints []point)
	dfs = func(tmpPoints []point) {
		if len(tmpPoints) == 1 {
			fmt.Print("O")
			return
		}
		// 第二部分头部为：
		mid := len(tmpPoints) / 2
		if step%2 == 0 {
			sort.Slice(tmpPoints, func(i, j int) bool {
				if tmpPoints[i].x < tmpPoints[j].x {
					return true
				} else if tmpPoints[i].x == tmpPoints[j].x {
					return tmpPoints[i].y < tmpPoints[j].y
				} else {
					return false
				}
			})
			if fpoint.x < tmpPoints[mid].x {
				onLeft = true
				fmt.Print("L")
			} else if fpoint.x == tmpPoints[mid].x {
				if fpoint.y < tmpPoints[mid].y {
					onLeft = true
					fmt.Print("L")
				} else {
					onLeft = false
					fmt.Print("R")
				}
			} else {
				onLeft = false
				fmt.Print("R")
			}
		} else {
			sort.Slice(tmpPoints, func(i, j int) bool {
				if tmpPoints[i].y < tmpPoints[j].y {
					return true
				} else if tmpPoints[i].y == tmpPoints[j].y {
					return tmpPoints[i].x < tmpPoints[j].x
				} else {
					return false
				}
			})
			if fpoint.y < tmpPoints[mid].y {
				onLeft = true
				fmt.Print("L")
			} else if fpoint.y == tmpPoints[mid].y {
				if fpoint.x < tmpPoints[mid].x {
					onLeft = true
					fmt.Print("L")
				} else {
					onLeft = false
					fmt.Print("R")
				}
			} else {
				onLeft = false
				fmt.Print("R")
			}
		}
		step++
		if onLeft {
			dfs(tmpPoints[:mid])
		} else {
			dfs(tmpPoints[mid:])
		}
	}
	dfs(points)
}

func deepCopy(s []point) []point {
	res := make([]point, len(s))
	copy(res, s)
	return res
}
