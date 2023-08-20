/*
 * @Author: liziwei01
 * @Date: 2023-08-19 23:53:20
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-19 23:53:45
 * @Description: file content
 */
package meituan

import "fmt"

func psum(str string) {
	sum := 0
	n := len(str)
	// get substrs
	for i := 2; i <= n; i++ {
		substrs := getSubstrs(str, i)
		for j := 0; j < len(substrs); j++ {
			sum += pw(substrs[j])
		}
	}

	fmt.Printf("%d\n", sum)
}

func pw(str string) int {
	sum := 0
	if len(str) == 1 {
		return 0
	}

	if len(str) == 2 {
		if str[0] == str[1] {
			return 1
		}
		return 0
	}

	for i := 2; i < len(str); i++ {
		if str[i-2] == str[i-1] && str[i-1] == str[i] {
			// 三数相同换中间
			sum++
			str = str[:i-1] + revStr(str[i-1]) + str[i:]
		} else if i+1 < len(str) {
			if str[i-2:i+2] == "1001" {
				sum += 2
				str = str[:i-2] + "0101" + str[i+3:]
			} else if str[i-2:i+2] == "0011" {
				sum += 2
				str = str[:i-1] + revStr(str[i-1]) + str[i:]
			}
		} else if str[i-2] == str[i-1] && str[i-1] != str[i] {
			temp1 := pw(str[i-1:]) + 1
			// 试试交换
			str = str[:i-1] + revStr(str[i-1]) + revStr(str[i]) + str[i+1:]
			temp2 := pw(str[i-1:]) + 2
			if temp1 < temp2 {
				return temp1
			} else {
				return temp2
			}

		} else if str[i-2] != str[i-1] && str[i-1] == str[i] {
			// 试试换前面
			temp1 := pw(str[i:])
			// 后俩相同换后面
			sum++
			str = str[:i] + revStr(str[i]) + str[i+1:]
			temp2 := pw(str[i-1:])
			if sum-1+2+temp1 < sum+temp2 {
				return sum - 1 + 2 + temp1
			} else {
				return sum + temp2
			}
		}
	}

	// for i := 2; i < len(str); i++ {
	// 	if str[i-2] == str[i-1] && str[i-1] == str[i] {
	// 		return 1 + pw(str[i:])
	// 	} else if str[i-2] == str[i-1] && str[i-1] != str[i] {
	// 		return 1 + pw(str[i-1:])
	// 	} else {
	// 		return pw(str[i-1:])
	// 	}
	// }

	return sum
}

func revStr(s byte) string {
	if s == '0' {
		return "1"
	}

	return "0"
}

func getSubstrs(str string, n int) []string {
	ret := make([]string, 0)
	for i := n; i <= len(str); i++ {
		ret = append(ret, str[i-n:i])
	}
	return ret
}
