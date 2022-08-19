/*
 * @Author: liziwei01
 * @Date: 2022-08-17 15:03:02
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-08-18 22:44:49
 * @Description: file content
 */
package countAndSay

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	num := countAndSay(n - 1)

	return say(num)
}

func say(num string) string {
	var res string
	var count int
	for i := 0; i < len(num); i++ {
		if i == 0 {
			count = 1
		} else if num[i] != num[i-1] {
			res += strconv.Itoa(count) + string(num[i-1])
			count = 1
		} else {
			count++
		}
	}

	res += strconv.Itoa(count) + string(num[len(num)-1])
	return res
}
