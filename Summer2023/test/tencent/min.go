/*
 * @Author: liziwei01
 * @Date: 2023-09-10 09:08:15
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-10 09:26:43
 * @Description: file content
 */
package tencent

import (
	"fmt"
	"math/big"
	"sort"
)

type bi struct {
	idx  int
	diff int64
}

func min() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	var tmp int64 = 0
	var sum = big.NewInt(0)
	arr := make([]int64, 0)
	bis := make([]bi, 0)
	for i := 0; i != n; i++ {
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
		bigTmp := big.NewInt(tmp)
		sum.Add(sum, bigTmp)
		lowestBit := tmp & -tmp
		modified := tmp ^ lowestBit
		bis = append(bis, bi{
			idx:  i,
			diff: tmp - modified,
		})
	}
	sort.Slice(bis, func(i, j int) bool {
		return bis[i].diff < bis[j].diff
	})

	for i := 0; i != k; i++ {
		sum.Sub(sum, big.NewInt(bis[len(bis)-1-i].diff))
	}

	fmt.Printf("%d", sum.Int64())
}
