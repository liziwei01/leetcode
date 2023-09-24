/*
 * @Author: liziwei01
 * @Date: 2023-09-23 02:08:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-23 03:00:42
 * @Description: file content
 */
package netease

import (
	"fmt"
	"sort"
)

func sorted() {
	t, n := 0, 0
	fmt.Scan(&t)
	for i := 0; i != t; i++ {
		fmt.Scan(&n)
		as := make([]int, n)
		for j := 0; j != n; j++ {
			fmt.Scan(&as[j])
		}
		if isIncrease(as, n) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isIncrease(as []int, l int) bool {
	// 如果长度是奇数，比如3,2,1. 则全组可以互换位置
	if isOdd(l) {
		return true
	}
	// 如果长度是偶数
	odds := make([]int, 0)
	evens := make([]int, 0)
	for i := 1; i != l+1; i++ {
		if isOdd(i) {
			odds = append(odds, as[i-1])
		} else {
			evens = append(evens, as[i-1])
		}
	}
	sort.Ints(odds)
	sort.Ints(evens)
	for i := 0; i != len(evens); i++ {
		if odds[i] > evens[i] {
			return false
		}
	}
	
	return true
}

func isOdd(n int) bool {
	return n%2 == 1
}
