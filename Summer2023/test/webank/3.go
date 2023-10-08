/*
 * @Author: liziwei01
 * @Date: 2023-10-08 07:53:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-08 08:16:54
 * @Description: file content
 */
package webank

import (
	"fmt"
	"sort"
)

// 两个英雄等级不同时，高的可+1，几个英雄能到达target等级？AC
func main() {
	// n 个英雄
	n := 0
	fmt.Scan(&n)

	us := make([]int, n)
	for i := 0; i != n; i++ {
		// 初始等级
		fmt.Scan(&us[i])
	}
	sort.Ints(us)
	target := 2147483647
	sum := 0
	exceed := 0
	// 0 1 2 3 4, n = 5, tar = 3, sum = 1, exceed = 3, nohope = 1
	for i := n - 1; i >= 0; i-- {
		if us[i] > target {
			continue
		} else if us[i] == target {
			sum++
		} else {
			exceed = i + 1
			break
		}
	}
	min := us[0]
	nohope := 0
	// 最低一档等级无望达成目标
	for i := 1; i < exceed; i++ {
		if us[i] == min {
			continue
		} else {
			nohope = i
			break
		}
	}
	sum += exceed - nohope
	fmt.Println(sum)
}