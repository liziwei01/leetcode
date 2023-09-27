/*
 * @Author: liziwei01
 * @Date: 2023-09-27 04:58:48
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-27 05:12:19
 * @Description: file content
 */
package shunfeng

import (
	"fmt"
	"sort"
)

func main() {
	N, M := 5, 100
	// fmt.Scan(&N, &M)
	As := make([]int, N)
	// for i := 0; i != N; i++ {
	// 	fmt.Scan(&As[i])
	// }
	As[0] = 82
	As[1] = 37
	As[2] = 43
	As[3] = 121
	As[4] = 55

	sort.Sort(sort.Reverse(sort.IntSlice(As)))
	i, j := 0, N-1
	usedI := false
	numofPeople, maxAbility := 0, 0
	numofGroup := 0
	for i <= j + 1 {
		if numofPeople * maxAbility < M {
			numofPeople++
			if !usedI {
				// if As[i] > maxAbility {
				// 	maxAbility = As[i]
				// }
				maxAbility = As[i]
				i++
				usedI = true
			} else {
				j--
			}
		} else {
			numofGroup++
			numofPeople, maxAbility = 0, 0
			usedI = false
		}
	}
	fmt.Print(numofGroup)
}
