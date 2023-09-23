/*
 * @Author: liziwei01
 * @Date: 2023-09-17 03:41:17
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-17 03:41:17
 * @Description: file content
 */

package aliyun

import (
	"fmt"
	"math/big"

	// "strings"
	"strconv"
)

func goodNum() {
	n := ""
	fmt.Scan(&n)
	l := len(n)
	if l <= 2 {
		fmt.Println(n)
		return
	}
	sum := big.NewInt(99)
	for i := 3; i < l-1; i++ {
		for j := 0; j <= i; j++ {
			a := big.NewInt(0).Binomial(int64(i), int64(j))
			sum.Add(sum, a.Mul(a, big.NewInt(100)))
		}
	}
	onezeroes := "1"
	for i := 0; i != l-1; i++ {
		onezeroes += "0"
	}
	start, _ := strconv.Atoi(onezeroes)
	end, _ := strconv.Atoi(n)
	// fmt.Println(start)
	// fmt.Println(end)
	for i := start; i <= end; i++ {
		if isGood(i) {
			sum.Add(sum, big.NewInt(1))
		}
	}
	fmt.Println(sum.Int64())
}

func isGood(n int) bool {
	s := strconv.Itoa(n)
	z, o, t, th, f, fi, si, se, e, ni := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	for i := 0; i != len(s); i++ {
		if s[i] == '0' {
			z = 1
		} else if s[i] == '1' {
			o = 1
		} else if s[i] == '2' {
			t = 1
		} else if s[i] == '3' {
			th = 1
		} else if s[i] == '4' {
			f = 1
		} else if s[i] == '5' {
			fi = 1
		} else if s[i] == '6' {
			si = 1
		} else if s[i] == '7' {
			se = 1
		} else if s[i] == '8' {
			e = 1
		} else if s[i] == '9' {
			ni = 1
		}
		if z+o+t+th+f+fi+si+se+e+ni > 2 {
			return false
		}
	}
	return true
}
