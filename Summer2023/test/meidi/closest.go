/*
 * @Author: liziwei01
 * @Date: 2023-09-28 04:24:17
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-28 04:55:10
 * @Description: file content
 */
package meidi

import (
	"fmt"
	"math/big"
)

func main() {
	n := 0
	fmt.Scan(&n)
	xs := make([]int, n)
	ys := make([]int, n)
	for i := 0; i != n; i++ {
		fmt.Scan(&xs[i])
	}
	for i := 0; i != n; i++ {
		fmt.Scan(&ys[i])
	}
	var ret float64 = 1<<63 - 1
	for i := 0; i != n; i++ {
		for j := i + 1; j != n; j++ {
			ret = min(ret, distance(xs[i], ys[i], xs[j], ys[j]))
		}
	}
	fmt.Printf("%.4f", ret)
}

func distance(x1, y1, x2, y2 int) float64 {
	xd2f := big.NewFloat(float64((x1 - x2) * (x1 - x2)))
	yd2f := big.NewFloat(float64((y1 - y2) * (y1 - y2)))
	xd2f.Add(xd2f, yd2f)
	df, _ := xd2f.Sqrt(xd2f).Float64()
	return df
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func abs(i float64) float64 {
	if i < 0 {
		return -i
	}
	return i
}
