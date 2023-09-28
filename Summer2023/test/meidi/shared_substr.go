/*
 * @Author: liziwei01
 * @Date: 2023-09-28 03:57:12
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-28 04:45:36
 * @Description: file content
 */
package meidi

import (
	"fmt"
	"strings"
)

func substr() {
	N, M := 0, 0
	N, M = 3, 4
	// fmt.Scan(&N, &M)
	S, T := "", ""
	S, T = "aaa", "aaaa"
	// fmt.Scan(&S, &T)
	if M > N {
		M, N = N, M
		S, T = T, S
	}
	max := 0
	for i := 0; i < M; i++ {
		for j := i+1; j <= M; j++ {
			Tij := T[i:j]
			if strings.Contains(S, Tij) {
				if max < j-i {
					max = j-i
				}
			}
		}
	}
	fmt.Println(max)
}
