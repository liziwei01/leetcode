/*
 * @Author: liziwei01
 * @Date: 2023-09-23 02:50:33
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-23 03:25:56
 * @Description: file content
 */
package netease

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func tree() {
	n, q := 0, 0
	fmt.Scan(&n, &q)
	m := make(map[int][]int)
	for i := 0; i != n-1; i++ {
		u, v := 0, 0
		fmt.Scan(&u, &v)
		m[u] = append(m[u], v)
		m[v] = append(m[v], u)
	}

	hasntVisited := make([]int, 0)
	minPath := 0
	path := make([]int, 0)
	pathNum := 1

	var dfs func(n int)
	dfs = func(n int) {
		if len(hasntVisited) == 0 {
			if pathNum < minPath {
				minPath = pathNum
			}
			return
		}
		for k, v := range hasntVisited {
			if slices.Contains(m[n], v) {
				hasntVisited = append(hasntVisited[:k], hasntVisited[k+1:]...)
				path = append(path, v)
				dfs(v)
				path = path[:len(path)-1]
				tmpTail := hasntVisited[k+1:]
				hasntVisited = append(hasntVisited[:k], v)
				hasntVisited = append(hasntVisited, tmpTail...)
			}
		}
	}

	for i := 0; i != q; i++ {
		m := 0
		as := make([]int, m)
		for j := 0; j != m; j++ {
			fmt.Scan(&as[j])
		}
		path = make([]int, 0)
		dfs(as[0])
		pathNum = 1
		for j := 0; j != len(path)-1; j++ {
			if len(path) > j+2 && path[j] == path[j+2] {
				pathNum++
			}
		}
		fmt.Println(pathNum)
	}
}
