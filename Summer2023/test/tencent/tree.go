package tencent

//import "fmt"
// import . "nc_tools"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 返回节点权值1个数比0的个数多一的路径数
 * @param root TreeNode类 权值为0和1的二叉树根节点
 * @return int整型
 */
func pathNumber(root *TreeNode) int {
	zeros, ones, pow := 0, 0, 0
	var dfs func(node *TreeNode)
	dfs = func(this *TreeNode) {
		if this == nil {
			return
		}
		if this.Val == 0 {
			zeros++
		} else {
			ones++
		}
		if this.Left == nil && this.Right == nil {
			if ones-zeros == 1 {
				pow++
			}
		}
		dfs(this.Left)
		dfs(this.Right)
		if this.Val == 0 {
			zeros--
		} else {
			ones--
		}
	}
	dfs(root)

	return pow
}
