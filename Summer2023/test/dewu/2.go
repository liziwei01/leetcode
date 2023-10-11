/*
 * @Author: liziwei01
 * @Date: 2023-10-09 22:23:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-09 22:48:09
 * @Description: file content
 */
package dewu

import (
	"fmt"
	"sort"
)

func ta() {
	N, M := 1000, 100000
	// N, M := 5, 5
	fmt.Scan(&N, &M)
	Ns := make([]int, N)
	// for i := 0; i != N; i++ {
	// 	fmt.Scan(&Ns[i])
	// }
	// Ns[0] = 1
	// Ns[1] = 3
	// Ns[2] = 2
	// Ns[3] = 1
	// Ns[4] = 1
	sort.Sort(sort.Reverse(sort.IntSlice(Ns)))
	// 找到的解的数字个数
	nums := make([]int, 0)
	num := 1
	var dfs func(sum, index int)
	dfs = func(sum, index int) {
		if sum == 0 {
			nums = append(nums, num)
			return
		}
		num++
		for i := index + 1; i < N; i++ {
			if Ns[i] <= sum {
				dfs(sum-Ns[i], i)
			}
		}
		num--
	}
	for i := 0; i < N; i++ {
		if Ns[i] <= M {
			dfs(M-Ns[i], i)
		}
	}
	if len(nums) == 0 {
		fmt.Println("No solution")
		return
	}
	sort.Ints(nums)
	fmt.Println(nums[0])
}
