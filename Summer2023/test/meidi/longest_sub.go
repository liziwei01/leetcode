/*
 * @Author: liziwei01
 * @Date: 2023-09-28 04:03:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-28 04:25:14
 * @Description: file content
 */
package meidi

import (
	"fmt"
)

func longest() {
	// N条绳子切K条
	N, K := 0, 0
	fmt.Scan(&N, &K)
	Ns := make([]float64, N)
	var maxString float64 = 0
	for i := 0; i != N; i++ {
		fmt.Scan(&Ns[i])
		if Ns[i] > maxString {
			maxString = Ns[i]
		}
	}
	// 剪枝
	if K == 1 {
		fmt.Print(maxString)
		return
	}
	Ns[0] = 8.02
	Ns[1] = 7.43
	Ns[2] = 4.57
	Ns[3] = 5.39

	floorLen := 0
	for i := 1; i != int(maxString); i++ {
		maxNumOfCutted := 0
		// 过一遍所有绳子
		for j := 0; j != N; j++ {
			maxNumOfCutted += cuttable(Ns[j], float64(i))
		}
		if maxNumOfCutted >= K {
			floorLen = i
		} else {
			break
		}
	}
	var i float64 = float64(floorLen)
	for i < float64(floorLen) + 1 {
		maxNumOfCutted := 0
		i += 0.01
		for j := 0; j != N; j++ {
			maxNumOfCutted += cuttable(Ns[j], i)
		}
		if maxNumOfCutted < K {
			i -= 0.01
			break
		}
	}
	fmt.Printf("%.2f", i)
}

// s能切出几条t
func cuttable(s, t float64) int {
	return int(s / t)
}
