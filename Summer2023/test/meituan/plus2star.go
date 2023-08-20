/*
 * @Author: liziwei01
 * @Date: 2023-08-19 23:51:43
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-19 23:52:21
 * @Description: file content
 */
package meituan

import "fmt"

func plus2star(n int, as []int) {
	// n := 0
	// fmt.Scan(&n)

	// as := make([]int, 0)
	// for i := 0; i != n; i++ {
	// 	a := 0
	// 	fmt.Scan(&a)
	// 	as = append(as, a)
	// }

	if n == 1 {
		fmt.Printf("%d\n", as[0])
		return
	}
	if n == 2 {
		fmt.Printf("%d\n", as[0]*as[1])
		return
	}

	biggestNum := as[0] * as[1]
	biggestIdx := 0
	for i := 1; i != n-1; i++ {
		if as[i]*as[i+1] > biggestNum {
			biggestIdx = i
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		if i != biggestIdx {
			sum += as[i]
		} else {
			sum += as[i] * as[i+1]
			i++
		}
	}

	fmt.Printf("%d\n", sum)
}
