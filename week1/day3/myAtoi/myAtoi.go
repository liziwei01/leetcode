/*
 * @Author: liziwei01
 * @Date: 2022-06-14 14:22:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-06-15 14:50:42
 * @Description: file content
 */
package myatoi

func myAtoi(s string) int {
	ret := make([]byte, 0)
	positive := true

	if len(s) == 0 {
		return 0
	}

	// 去除先导空格
	for k, v := range s {
		if isSpace(v) {
			continue
		} else {
			s = s[k:]
			break
		}
	}

	// 检查正负号
	if s[0] == '+' {
		positive = true
		s = s[1:]
	} else if s[0] == '-' {
		positive = false
		s = s[1:]
	}

	// 读入数字
	for _, v := range s {
		if isInt(v) {
			if v == '0' && len(ret) == 0 {
				continue
			}
			ret = append(ret, byte(v))
		} else {
			break
		}
	}

	// 截断超长数字
	if ExceedIntRange(ret, !positive) {
		if positive {
			return 2147483647
		} else {
			return -2147483648
		}
	}

	if positive {
		return Int(ret)
	} else {
		return -Int(ret)
	}
}

func isSpace(s rune) bool {
	return s == ' '
}

func isInt(s rune) bool {
	return s >= '0' && s <= '9'
}

func ExceedIntRange(xAbsByte []byte, isMinus bool) bool {
	length := len(xAbsByte)
	benchmark := []byte("2147483647")
	if length > 10 {
		return true
	}
	if length < 10 {
		return false
	}
	if isMinus {
		benchmark = []byte("2147483648")
	} else {
		benchmark = []byte("2147483647")
	}
	for i := 0; i < len(benchmark); i++ {
		x := xAbsByte[i]
		b := benchmark[i]
		if x > b {
			return true
		} else if x == b {
			continue
		} else {
			return false
		}
	}
	return false
}

func Int(b []byte) int {
	var result int
	for i := 0; i < len(b); i++ {
		result = result*10 + int(b[i]-'0')
	}
	return result
}
