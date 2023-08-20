/*
 * @Author: liziwei01
 * @Date: 2023-08-19 00:03:12
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-19 00:03:12
 * @Description: file content
 */
/*
 * @Author: liziwei01
 * @Date: 2023-07-17 14:00:02
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-19 00:02:05
 * @Description: file content
 */
package jd

import (
	"fmt"
)

func bd() {
	n := 0
	fmt.Scan(&n)
	s := make([]string, 0)

	for i := 0; i != n; i++ {
		for j := 0; j != 3; j++ {
			str := ""
			fmt.Scan(&str)
			s = append(s, str)
		}
	}

	for i := 0; i != n; i++ {
		k := kou(s[3*i : 3*i+3])
		y := yukari(s[3*i : 3*i+3])
		if k && y {
			fmt.Printf("drawk\n")
		} else if k {
			fmt.Printf("kou\n")
		} else if y {
			fmt.Printf("yukari\n")
		} else {
			fmt.Printf("draw\n")
		}
	}

	return
}

// black win
func kou(s []string) bool {
	return win(s, "*o*")
}

// white win
func yukari(s []string) bool {
	return win(s, "o*o")
}

func win(s []string, win string) bool {
	fmt.Print(s)
	for i := 0; i != 3; i++ {
		str := s[i]
		if str == win {
			return true
		}
	}
	for i := 0; i != 3; i++ {
		str := ""
		str += string(s[0][i])
		str += string(s[1][i])
		str += string(s[2][i])
		if str == win {
			return true
		}
	}
	return false
}
