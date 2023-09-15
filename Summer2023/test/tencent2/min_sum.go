/*
 * @Author: liziwei01
 * @Date: 2023-09-15 08:12:33
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-15 09:43:32
 * @Description: file content
 */
package tencent2

import (
	"fmt"
	"math/big"
	"sort"
)

func minSum() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	as := make([]int, n)
	sum := big.NewInt(0)
	for i := 0; i != n; i++ {
		fmt.Scan(&as[i])
		sum.Add(sum, big.NewInt(int64(as[i])))
	}
	sort.Ints(as)
	idx, mid := 0, n/4*2
	for i := 0; i != k; i++ {
		sum.Add(sum, big.NewInt(int64(operater(as, idx))))
		if as[idx] > as[mid] && idx < n/4*2 {
			idx = (idx + 1 ) % n
			mid = (idx+n/4*2) % n
		} else {
			as = append(as[idx:], as[:idx]...)
			firstItemSelectionSort(as)
			idx, mid = 0, n/4*2
		}
	}
	fmt.Print(sum.Int64())
}

func operater(data []int, i int) int {
	o := data[i]
	if isEven(data[i]) {
		data[i] = data[i]*2 + 1
	} else {
		data[i] = data[i] * 2
	}
	return data[i] - o
}

func isEven(a int) bool {
	return a%2 == 0
}

func firstItemSelectionSort(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			data[i], data[i-1] = data[i-1], data[i]
		} else {
			return
		}
	}
}
