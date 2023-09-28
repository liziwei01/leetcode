/*
 * @Author: liziwei01
 * @Date: 2023-09-28 07:57:15
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-28 08:07:05
 * @Description: file content
 */
package didi

func node() {
	n := 3
	// n, m := 0, 0
	// fmt.Scan(&n, &m)
	parent := make(map[int]int)
	childs := make(map[int][]int)
	for i := 1; i <= n; i++ {
		childs[i] = make([]int, 0)
	}
	// for i := 2; i <= n; i++ {
	// 	p := 0
	// 	fmt.Scan(&p)
	// 	parent[i] = p
	// 	childs[p] = append(childs[p], i)
	// }
	p := 1
	ii := 2
	parent[ii] = p
	childs[p] = append(childs[p], ii)
	ii++
	parent[ii] = p
	childs[p] = append(childs[p], ii)
	
	S(parent, childs, 2)

	// for i := 1; i <= m; i++ {
	// 	x := 0
	// 	fmt.Scan(&x)
	// 	distance := S(parent, childs, x)
	// 	fmt.Printf("%d ", distance)
	// }
}

func S(parent map[int]int, childs map[int][]int, x int) int {
	sum, distance := 0, 0
	visited := make(map[int]bool)
	var dfs func(node int)
	dfs = func(node int) {
		// 不能往回走
		if visited[node] {
			return
		}
		// 加上到此节点的距离
		sum += distance
		
		visited[node] = true
		distance++
		// 过一遍子节点
		for i := 0; i != len(childs[node]); i++ {
			dfs(childs[node][i])
		}
		// 有父节点尝试走一下父节点
		if v, ok := parent[node]; ok {
			dfs(v)
		}
		distance--
	}
	dfs(x)
	return sum
}
