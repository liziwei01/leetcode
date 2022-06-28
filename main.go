/*
 * @Author: liziwei01
 * @Date: 2022-06-14 13:37:41
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-28 15:02:13
 * @Description: file content
 */
package main

import "fmt"

func main() {
	a := &[]string{"a", "b", "c"}
	b := *a

	(*a)[0] = "d"
	fmt.Println(a)
	fmt.Println(b)
}
