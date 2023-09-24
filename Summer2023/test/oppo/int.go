/*
 * @Author: liziwei01
 * @Date: 2023-09-23 07:08:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-23 07:13:19
 * @Description: file content
 */
package oppo

import (
	"fmt"
)

func intsss() {
	// 两个区间，有多少个mid=(l+r)/2
	var l1, r1, l2, r2 int64 = 0, 0, 0, 0
	fmt.Scan(&l1, &r1, &l2, &r2)
	l := (l1 + l2) / 2
	r := (r1 + r2) / 2
	fmt.Print(r - l + 1)
}
