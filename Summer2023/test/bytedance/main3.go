/*
 * @Author: liziwei01
 * @Date: 2023-08-26 23:38:33
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-26 23:55:03
 * @Description: file content
 */
package bytedance

import (
	"fmt"
	"strconv"
)

func main3() {
	n := 3
	pSum := 0
	min, _ := strconv.Atoi("1" + zeros(n-1))
	max, _ := strconv.Atoi(nines(n))
	for i := min; i <= max; i++ {
		oddSum, evenSum := Sums(i)
		pSum = pSum + (oddSum * evenSum)
	}
	fmt.Print(pSum)
}

func Sums(a int) (int, int) {
	oddSum := 0
	evenSum := 0
	s := fmt.Sprintf("%d", a)
	i := 0
	for i < len(s) {
		o, _ := strconv.Atoi(s[i : i+1])
		oddSum += o
		i++
		if i == len(s) {
			break
		}
		e, _ := strconv.Atoi(s[i : i+1])
		evenSum += e
		i++
	}
	return oddSum, evenSum
}

func zeros(a int) string {
	z := ""
	for i := 0; i != a; i++ {
		z = z + "0"
	}
	return z
}

func nines(a int) string {
	z := ""
	for i := 0; i != a; i++ {
		z = z + "9"
	}
	return z
}
