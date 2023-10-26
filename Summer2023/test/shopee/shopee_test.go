/*
 * @Author: liziwei01
 * @Date: 2023-10-26 06:52:55
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-26 07:59:42
 * @Description: file content
 */
package shopee

import "testing"

func TestXxx(t *testing.T) {
	// {5,3,4,2,#,7,1}
	// tree := &TreeNode{
	// 	Val: 5,
	// 	Left: &TreeNode{
	// 		Val: 3,
	// 		Left: &TreeNode{
	// 			Val: 2,
	// 		},
	// 		Right: nil,
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 4,
	// 		Left: &TreeNode{
	// 			Val: 7,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 1,
	// 		},
	// 	},
	// }
	// t.Log(getMostGold(tree))

	t.Log(numSimilarGroups([]string{"abcd", "cbad", "bcad", "dabc"}))
}
