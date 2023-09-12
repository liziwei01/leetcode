/*
 * @Author: liziwei01
 * @Date: 2023-09-12 14:16:35
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-12 14:16:40
 * @Description: file content
 */
package channel

import (
	"fmt"
	"time"
)

func rangeExample() {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}
