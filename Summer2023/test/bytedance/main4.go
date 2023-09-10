/*
 * @Author: liziwei01
 * @Date: 2023-08-26 23:38:50
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-26 23:38:54
 * @Description: file content
 */
package bytedance

import (
    "fmt"
)

func main4() {
    l := 0
    r := 0
    fmt.Scan(&l, &r)
    sum := 0
    for i := l; i <= r; i++ {
        sum += f(i)
    }
    fmt.Print(sum)
}

func f(i int) int {
    m := 0
    for i != 0 {
        m = max(i%10, m)
        i /= 10
    }
    return m
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}