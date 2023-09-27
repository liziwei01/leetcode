/*
 * @Author: liziwei01
 * @Date: 2023-09-27 06:37:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-27 06:37:32
 * @Description: file content
 */
package zhangyue

import (
	"fmt"
	"sort"
	"strings"
)

type numforSort struct {
	i   string
	num int
}

func m1() {
	nums := ""
	k := 0
	fmt.Scan(&nums, &k)
	// fmt.Println(nums)
	// fmt.Println(k)
	numSlice := strings.Split(nums, ",")
	intMap := make(map[string]int)
	for i := 0; i != len(numSlice); i++ {
		intMap[numSlice[i]]++
	}
	intSlice := make([]numforSort, 0)
	for key, val := range intMap {
		intSlice = append(intSlice, numforSort{
			i:   key,
			num: val,
		})
	}
	sort.Slice(intSlice, func(i, j int) bool { return intSlice[i].num > intSlice[j].num })
	for i := 0; i < k-1; i++ {
		fmt.Printf("%s,", intSlice[i].i)
	}
	fmt.Printf("%s", intSlice[k-1].i)
}
