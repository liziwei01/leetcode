/*
 * @Author: liziwei01
 * @Date: 2023-08-19 23:54:43
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-19 23:54:44
 * @Description: file content
 */
package meituan

import (
	"fmt"
	"math/big"
)

func redistribute(n, sum int) {
	ret := combNum(sum, n)
	ret = mod(ret, 10*10*10*10*10*10*10*10*10+7)
	fmt.Printf("%d\n", ret)
}

func combNum(sum, n int) int {
	num := sum - n + 1
	start := sum - n
	for i := 0; i != n; i++ {
		num *= start
		start--
	}
	return num
}

// n 选 k
func comb(n, k int) int {
	return int(new(big.Int).Binomial(int64(n), int64(k)).Int64())
}

func mod(a, b int) int {
	if a > b {
		return mod(a-b, b)
	}
	return a
}
