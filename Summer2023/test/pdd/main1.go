/*
 * @Author: liziwei01
 * @Date: 2023-08-22 08:20:32
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-22 08:33:16
 * @Description: file content
 */
package pdd

import (
	"fmt"
)

func main1(T, K int, S string) {
	// T := 0
	// fmt.Scan(&T)

	// for i := 0; i != T; i++ {
	// S := ""
	// K := 0
	// fmt.Scan(&S, &K)

	bs := []byte(S)

	length := len(S)
	d := distance(bs, length)
	if K >= d {
		fmt.Print("YES\n")
	} else {
		fmt.Print("NO\n")
	}
	// }
	return
}

func distance(bs []byte, length int) int {
	d := 0

	for i := 0; i != length/2; i++ {
		if bs[i] != bs[length-i-1] {
			d++
		}
	}

	return d
}
