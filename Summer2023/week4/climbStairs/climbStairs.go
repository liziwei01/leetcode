/*
 * @Author: liziwei01
 * @Date: 2023-07-09 15:38:56
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-09 15:45:15
 * @Description: file content
 */
package climbstairs

import "math/big"

func climbStairs(n int) int {
	sum := 0
	maxTwos := n / 2
	for i := 0; i != maxTwos+1; i++ {
		ones := n - i*2
		twos := i
		sum += comb(ones+twos, ones)
	}
	return sum
}

// n 选 k
func comb(n, k int) int {
	return int(new(big.Int).Binomial(int64(n), int64(k)).Int64())
}
