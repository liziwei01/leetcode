/*
 * @Author: liziwei01
 * @Date: 2023-10-26 07:35:50
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-26 07:37:10
 * @Description: file content
 */
package shopee

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 * 
 * 
 * @param str string字符串  
 * @param n int整型  
 * @return string字符串
*/
func moveSubStr(str string, n int) string {
	n--
	return str[n:] + str[:n]
}
