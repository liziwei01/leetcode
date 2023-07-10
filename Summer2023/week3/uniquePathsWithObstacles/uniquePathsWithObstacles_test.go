/*
 * @Author: liziwei01
 * @Date: 2023-07-08 16:13:37
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-08 18:39:31
 * @Description: file content
 */
package uniquepathswithobstacles

import "testing"

func TestXxx(t *testing.T) {
	// t.Log(uniquePathsWithObstacles([][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}))
	t.Log(uniquePathsWithObstacles([][]int{{0, 0, 0, 1}, {0, 0, 1, 0}, {0, 1, 0, 0}}))
}
