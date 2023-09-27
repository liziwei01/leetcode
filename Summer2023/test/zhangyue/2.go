/*
 * @Author: liziwei01
 * @Date: 2023-09-27 06:37:37
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-27 06:37:38
 * @Description: file content
 */

package zhangyue

import (
	"fmt"
	"sort"
	"strconv"
)

type intforSort struct {
	i        int
	isASlice bool
}

func m2() {
	a := ""
	aSlice := make([]int, 0)
	bSlice := make([]int, 0)
	usingASlice := true
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			if !usingASlice {
				intVer, _ := strconv.Atoi(a[:len(a)-1])
				bSlice = append(bSlice, intVer)
			} else if usingASlice && a[len(a)-1] == ',' {
				intVer, _ := strconv.Atoi(a[:len(a)-1])
				aSlice = append(aSlice, intVer)
			} else {
				usingASlice = false
				intVer, _ := strconv.Atoi(a)
				aSlice = append(aSlice, intVer)
			}
		}
	}
	intVer, _ := strconv.Atoi(a)
	bSlice[len(bSlice)-1] = intVer
	// fmt.Println(aSlice)
	// fmt.Println(bSlice)
	abSlice := make([]intforSort, 0)
	for i := 0; i != len(aSlice); i++ {
		abSlice = append(abSlice, intforSort{
			i:        aSlice[i],
			isASlice: true,
		})
	}
	for i := 0; i != len(bSlice); i++ {
		abSlice = append(abSlice, intforSort{
			i:        bSlice[i],
			isASlice: false,
		})
	}
	sort.Slice(abSlice, func(i, j int) bool { return abSlice[i].i < abSlice[j].i })
	min := 1<<32 - 1
	for i := 0; i < len(abSlice)-1; i++ {
		if (abSlice[i].isASlice == true && abSlice[i+1].isASlice == false) || (abSlice[i].isASlice == false && abSlice[i+1].isASlice == true) {
			absolute := abs(abSlice[i].i - abSlice[i+1].i)
			if absolute < min {
				min = absolute
			}
		}
	}
	fmt.Print(min)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
