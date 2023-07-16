/*
 * @Author: liziwei01
 * @Date: 2023-07-15 02:30:43
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-15 02:42:23
 * @Description: file content
 */
package exist

import "testing"

func TestXxx(t *testing.T) {
	t.Log(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCESEEEFS"))
}
