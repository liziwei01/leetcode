/*
 * @Author: liziwei01
 * @Date: 2023-09-10 09:30:38
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-10 09:31:41
 * @Description: file content
 */
package tencent

import "fmt"

func verify() {
	T, n, k := 0, 0, 0
	fmt.Scan(&T)
	for i := 0; i != T; i++ {
		fmt.Scan(&n, &k)
	}
}