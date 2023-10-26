/*
 * @Author: liziwei01
 * @Date: 2023-10-26 07:38:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-26 07:51:17
 * @Description: file content
 */
package shopee

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ValSum struct {
	Sum int
	Route []int
}

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 *
 * @param nodes TreeNode类
 * @return int整型一维数组
 */
func getMostGold(nodes *TreeNode) []int {
	vals := make([]ValSum, 0)
	val := ValSum{}
	var dfs func(nodes *TreeNode)
	dfs = func(nodes *TreeNode) {
		if nodes == nil {
			return
		}
		val.Sum += nodes.Val
		val.Route = append(val.Route, nodes.Val)
		// 到达叶子节点
		if nodes.Left == nil && nodes.Right == nil {
			vals = append(vals, ValSum{
				Sum:   val.Sum,
				Route: append([]int{}, val.Route...),
			})
			val.Sum -= nodes.Val
			val.Route = val.Route[:len(val.Route)-1]
			return
		}
		dfs(nodes.Left)
		dfs(nodes.Right)
		val.Sum -= nodes.Val
		val.Route = val.Route[:len(val.Route)-1]
	}
	dfs(nodes)
	sort.Slice(vals, func(i, j int) bool {
		return vals[i].Sum > vals[j].Sum
	})
	return vals[0].Route
}
