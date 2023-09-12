/*
 * @Author: liziwei01
 * @Date: 2023-09-12 14:17:33
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-12 14:20:16
 * @Description: file content
 */
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
