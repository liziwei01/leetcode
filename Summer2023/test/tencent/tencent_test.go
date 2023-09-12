/*
 * @Author: liziwei01
 * @Date: 2023-09-10 07:53:36
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-10 09:35:54
 * @Description: file content
 */
package tencent

import "testing"

func TestTecent(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val: 0,
		Left: nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val: 0,
		Left: nil,
		Right: nil,
	}
	root.Left.Left = &TreeNode{
		Val: 1,
		Left: nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val: 0,
		Left: nil,
		Right: nil,
	}
	root.Right.Right = &TreeNode{
		Val: 1,
		Left: nil,
		Right: nil,
	}
	t.Log(pathNumber(root))
}
