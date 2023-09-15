/*
 * @Author: liziwei01
 * @Date: 2023-09-15 08:43:01
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-15 09:53:45
 * @Description: file content
 */
package tencent2

import (
	"fmt"
	"sort"
)

type position struct {
	originalPosition int
	position         int
	originalSeq      int
	seq              int
}

func game() {
	n, t := 0, 0
	fmt.Scan(&n, &t)
	ps := make([]position, n)
	vs := make([]int, n)
	for i := 0; i != n; i++ {
		fmt.Scan(&ps[i].originalPosition)
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].originalPosition > ps[j].originalPosition })
	ps[0].originalSeq = 1
	for i := 1; i < n; i-- {
		if ps[i].originalPosition == ps[i-1].originalPosition {
			ps[i].originalSeq = ps[i-1].originalSeq
		} else {
			ps[i].originalSeq = n - i
		}
	}
	for i := 0; i != n; i++ {
		fmt.Scan(&vs[i])
	}
	for i := 0; i != n; i++ {
		ps[i].position = ps[i].originalPosition + vs[i]*t
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].position > ps[j].position })
	sum := 0
	ps[0].seq = 1
	for i := 1; i < n; i++ {
		if ps[i].position == ps[i-1].position {
			ps[i].seq = ps[i-1].seq
		} else {
			ps[i].seq = n - i
		}
		if ps[i].seq < ps[i].originalSeq {
			sum++
		}
	}
	fmt.Print(sum)
}
