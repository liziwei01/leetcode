/*
 * @Author: liziwei01
 * @Date: 2023-09-23 02:31:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-23 03:08:19
 * @Description: file content
 */
package netease

import (
	"fmt"
	"math/big"
)

func sum() {
	n := 3
	fmt.Scan(&n)
	as := make([]float64, n)
	// as[0] = 1
	// as[1] = 2
	// as[2] = 3
	sum := big.NewFloat(0)
	for i := 0; i != n; i++ {
		fmt.Scan(&as[i])
		sum.Add(sum, big.NewFloat(as[i]))
	}
	sum.Mul(sum, big.NewFloat(float64((n-1)*n+1)))
	sum.Quo(sum, big.NewFloat(float64(n)))
	// sum64, aas := sum.Int64()
	// newSum := big.NewInt(sum64)
	for sum.Cmp(big.NewFloat(1000000007)) == 1 {
		sum.Add(sum, big.NewFloat(-1000000007))
	}
	fmt.Print(sum.Float64())
}
