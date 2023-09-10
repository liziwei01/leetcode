/*
 * @Author: liziwei01
 * @Date: 2023-08-27 09:29:25
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-27 09:52:07
 * @Description: file content
 */
package mhy

import (
	"fmt"
	"math/big"
	"strings"
)

func main3() {
	n := 0
	tree := ""
	fmt.Scan(&n, &tree)
	edges1 := make(map[int]int)
	edges2 := make(map[int]int)
	for i := 0; i != n-1; i++ {
		u, v := 0, 0
		fmt.Scan(&u, &v)
		edges1[u-1] = v - 1
		edges2[v-1] = u - 1
	}

	powSum := big.NewInt(0)

	for i := 0; i != n; i++ {
		for j := 1; j != n; i++ {
			p := path(tree, i, j, edges1, edges2)
			idx := 0
			for idx != -1 {
				idx = strings.Index(p[idx+3:], "mhy")
				if idx != -1 {
					powSum.Add(powSum, big.NewInt(1))
				}
			}

		}
	}
}

func path(tree string, i, j int, edges1, edges2 map[int]int) string {
	var dfs func(now int)
	visited := make(map[int]bool)
	p := ""
	dfs = func(now int) {
		if now == j {
			return
		}
		next1 := edges1[now]
		next2 := edges2[now]
		if visited[next1] && visited[next2] {
			// 走到头了也没找到j，这条路走不通
			return
		}
		if !visited[next1] {
			visited[next1] = true
			p += tree[now : now+1]
			dfs(next1)
			visited[next1] = false
			p = p[:len(p)-1]
		}
		if !visited[next2] {
			visited[next2] = true
			p += tree[now : now+1]
			dfs(next2)
			visited[next2] = false
			p = p[:len(p)-1]
		}
	}
	dfs(i)
	return p
}
