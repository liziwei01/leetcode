/*
 * @Author: liziwei01
 * @Date: 2022-06-14 12:30:53
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-14 14:22:40
 * @Description: file content
 */
package zconvert

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	length := len(s)
	result := make([]string, numRows)

	for i := 0; i < length; i++ {
		if i%(numRows*2-2) < numRows {
			result[i%(numRows*2-2)] += string(s[i])
		} else {
			result[numRows*2-2-i%(numRows*2-2)] += string(s[i])
		}
	}

	var resultStr string
	for _, v := range result {
		resultStr += v
	}

	return resultStr
}
