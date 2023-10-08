/*
 * @Author: liziwei01
 * @Date: 2023-10-08 07:37:11
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-08 08:15:38
 * @Description: file content
 */
package webank

import (
	"fmt"
	"sort"
	"strconv"
)

// CD之间距离为0，从A到B最短距离？WA
func roadq() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	A, B, C, D := 0, 0, 0, 0
	fmt.Scan(&A, &B, &C, &D)
	if C > D {
		C, D = D, C
	}

	roadMap := make(map[string]int)
	u, v, c := 0, 0, 0
	for i := 0; i != m; i++ {
		fmt.Scan(&u, &v, &c)
		if u > v {
			u, v = v, u
		}
		road := strconv.Itoa(u) + strconv.Itoa(v)
		roadMap[road] = c
		if u == C && v == D {
			roadMap[road] = 0
		}
	}
	distances := make([]int, 0)
	distance := 0
	visited := make(map[int]bool)
	var dfs func(town int)
	dfs = func(town int) {
		// 不走回头路
		if visited[town] {
			return
		}
		// 到达目标地点
		if town == B {
			distances = append(distances, distance)
			return
		}
		// 记录已访问地点
		visited[town] = true
		// 遍历可能的下一个地点
		for i := 1; i <= n; i++ {
			tmpA, tmpB := town, i
			if tmpA > tmpB {
				tmpA, tmpB = tmpB, tmpA
			}
			road := strconv.Itoa(tmpA) + strconv.Itoa(tmpB)
			// 有路就走
			if d, ok := roadMap[road]; ok {
				distance += d
				dfs(i)
				distance -= d
			}
		}
		visited[town] = false
	}
	dfs(A)
	sort.Ints(distances)
	fmt.Println(distances[0])
}
