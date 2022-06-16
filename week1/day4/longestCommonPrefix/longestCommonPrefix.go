/*
 * @Author: liziwei01
 * @Date: 2022-06-16 08:00:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-16 08:13:48
 * @Description: file content
 */
package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	length := len(strs)

	if length == 0 {
		return ""
	}

	for i := 0; i < length; i++ {
		if len(strs[i]) == 0 {
			return ""
		}
	}

	commonPrefix := strs[0][0:1]
	strs[0] = strs[0][1:]

	for i := 1; i < length; i++ {
		if commonPrefix == strs[i][0:1] {
			strs[i] = strs[i][1:]
		} else {
			return ""
		}
	}
	return commonPrefix + longestCommonPrefix(strs)
}
