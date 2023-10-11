/*
 * @Author: liziwei01
 * @Date: 2023-10-09 22:17:04
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-09 22:46:43
 * @Description: file content
 */
package dewu

import (
	"fmt"
)

func pa() {
	n, x := 0, 0
	fmt.Scan(&n, &x)
	s := ""
	fmt.Scan(&s)
	for i := 0; i != n-x+1; i++ {
		if isPalindrome(s[i : i+x]) {
			fmt.Println(1)
			return
		}
	}
	fmt.Println(0)
}

func isPalindrome(sub string) bool {
	if len(sub) == 0 || len(sub) == 1 {
		return true
	}
	if sub[0] == sub[len(sub)-1] {
		return isPalindrome(sub[1 : len(sub)-1])
	}
	return false
}
