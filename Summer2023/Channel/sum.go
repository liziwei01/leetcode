/*
 * @Author: liziwei01
 * @Date: 2023-09-12 14:16:11
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-12 14:17:17
 * @Description: file content
 */
package channel

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sumExample() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}
