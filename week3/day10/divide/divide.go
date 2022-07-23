/*
 * @Author: liziwei01
 * @Date: 2022-07-21 18:17:28
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-21 18:46:28
 * @Description: file content
 */
package divide

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	i := 0
	j := 0

	absDividend := Abs(dividend)
	absDivisor := Abs(divisor)

	if absDivisor == 1 || absDivisor == -1 {
		return CheckOverFlow(dividend, divisor, absDividend)
	}

	multiN := 0
	OriginAbsDivisor := absDivisor
	for absDividend > (absDivisor << 1) {
		multiN += 1
		absDivisor <<= 1
	}

	for absDividend >= absDivisor {
		i++
		absDividend -= absDivisor
	}

	for absDividend >= OriginAbsDivisor {
		j++
		absDividend -= OriginAbsDivisor
	}

	return CheckOverFlow(dividend, divisor, i<<multiN+j)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func SameSign(x int, y int) bool {
	return (x > 0 && y > 0) || (x < 0 && y < 0)
}

func CheckOverFlow(dividend int, divisor int, i int) int {
	if !SameSign(dividend, divisor) {
		i = -i

		if i < -2147483648 {
			return 2147483647
		}
	}

	if i > 2147483647 {
		return 2147483647
	}

	return i
}
