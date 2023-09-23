/*
 * @Author: liziwei01
 * @Date: 2023-09-17 03:40:37
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-17 03:40:38
 * @Description: file content
 */
package aliyun

import (
	"fmt"
	// "sort"
)

//	type edge struct {
//	    u int
//	    v int
//	}
type edge []int

func tree6() {
	n, u, v := 0, 0, 0
	fmt.Scan(&n)
	powers := make([]int, n)
	for i := 0; i != n; i++ {
		fmt.Scan(&powers[i])
	}
	// edges := make([]edge, n-1)
	edges := make(map[int]edge)
	for i := 0; i != n-1; i++ {
		fmt.Scan(&u, &v)
		if u > v {
			u, v = v, u
		}
		// edges[i].u = u
		// edges[i].v = v
		edges[u] = append(edges[u], v)
	}
	// sort.Slice(edges, func(i, j int) bool {
	//     if edges[i].u < edges[j].u {
	//         return true
	//     } else if edges[i].u == edges[j].u {
	//         return edges[i].v < edges[j].v
	//     }
	//     return false
	//     })
	var dfs func(node int)
	dfs = func(node int) {
		// if edges[node] == 0 {
		return
		// }
	}
	dfs(0)
}
