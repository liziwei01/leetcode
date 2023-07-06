/*
 * @Author: liziwei01
 * @Date: 2023-06-06 18:26:26
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-08 23:54:00
 * @Description: file content
 */
package multiply

import "fmt"

func multiply(num1 string, num2 string) string {
	l := New(num1)
	r := New(num2)
	l.Mul(r)
	return l.String()
}

type strInt struct {
	// 临时值，初始为等于constant的占位符
	TempVal      string
	// 溢出位
	Overflow int
	// 指向某一位数的指针
	Idx      int
	// 数值本体
	Constant string
}

func New(str string) *strInt {
	s := &strInt{
		TempVal:      str,
		Overflow: 0,
		Idx:      0,
		Constant: str,
	}
	s.Init()
	return s
}

func (l *strInt) Mul(r *strInt) {
	l.Order(r)
	l.Init()
	r.Init()
	l.mul(r)
	l.Constant = l.TempVal
	l.DeZeroes()
}

func (l *strInt) mul(r *strInt) {
	ans := New("")
	for i := r.Len() - 1; i >= 0; i-- {
		mult := r.getMultiplier(i)
		l.multOnce(mult)
		adder := New(l.TempVal + Zeros(r.Len()-i-1))
		ans.Add(adder)
	}
	l.TempVal = ans.TempVal
}

// 只乘以一位数
func (l *strInt) multOnce(r *strInt) {
	if l.IsEmpty() || r.IsEmpty() {
		l.Init()
		r.Init()
		return
	} else {
		ans := l.IntatIdx()*r.IntatIdx() + l.Overflow
		l.ReplaceatIdx(ans)
		l.Overflow = ans / 10
		r.Overflow = 0
		if l.Idx < 0 {
			l.Init()
			r.Init()
			return
		}
		l.Idx--
		l.multOnce(r)
	}
}

func (l *strInt) Add(r *strInt) {
	l.Order(r)
	l.Init()
	r.Init()
	l.add(r)
	l.Constant = l.TempVal
	l.DeZeroes()
}

func (l *strInt) add(r *strInt) {
	if len(l.Constant) == 0 {
		l.TempVal = r.Constant
		return
	}
	if l.IsEmpty() && r.IsEmpty() {
		return
	} else {
		ans := l.IntatIdx() + r.IntatIdx() + l.Overflow
		l.ReplaceatIdx(ans)
		l.Overflow = ans / 10
		r.Overflow = 0
		l.Idx--
		r.Idx--
		l.add(r)
	}
}

func (l *strInt) Init() {
	l.Idx = l.Len() - 1
	if l.Len() == 0 {
		l.Idx = 0
	}
}

func (l *strInt) SliceBound(i, j int) (int, int) {
	if i >= j || j <= 0 || i >= l.Len() {
		return 0, 0
	}
	if i >= 0 && j <= l.Len() {
		return i, j
	} else if i < 0 && j <= l.Len() {
		return 0, j
	} else if i >= 0 && j > l.Len() {
		return i, l.Len()
	} else if i < 0 && j > l.Len() {
		return 0, l.Len()
	}
	return 0, 0
}

func (l *strInt) Len() int {
	return len(l.Constant)
}

// 只换一位数
func (l *strInt) ReplaceatIdx(i int) {
	l.Replaceat(i, l.Idx)
}

func (l *strInt) Replaceat(i, idx int) {
	if i >= 10 {
		i = i % 10
	}
	l.replace(i, idx)
}

func (l *strInt) replace(i, idx int) {
	if idx < 0 {
		l.TempVal = fmt.Sprint(i) + l.TempVal
		return
	}
	bound1l, bound1r := l.SliceBound(0, idx)
	bound2l, bound2r := l.SliceBound(idx+1, l.Len())
	l.TempVal = l.TempVal[bound1l:bound1r] + fmt.Sprint(i) + l.TempVal[bound2l:bound2r]
}

func (l *strInt) getMultiplier(idx int) *strInt {
	if idx < 0 || idx >= l.Len() {
		return &strInt{}
	}

	r := &strInt{
		Constant: l.TempVal[idx : idx+1],
	}
	r.Init()

	return r
}

func Zeros(len int) string {
	ret := ""
	for i := 0; i != len; i++ {
		ret = ret + "0"
	}

	return ret
}

// 最多允许在前面添加一位数
// 1 没有初始化 或
// 2 没有前一位数要加且指针指向前一位 或
// 3 指针溢出
// 即定义为空
func (l *strInt) IsEmpty() bool {
	return len(l.Constant) == 0 || (l.Overflow == 0 && l.Idx == -1) || l.Idx <= -2
}

func (l *strInt) Order(r *strInt) {
	if l.Len() < r.Len() {
		*l, *r = *r, *l
	}
}

func (l *strInt) Byteat(idx int) byte {
	return l.Constant[idx]
}

func (l *strInt) Intat(idx int) int {
	if l.Len() == 0 || idx < 0 || idx >= l.Len() {
		return 0
	}
	return int(l.Byteat(idx) - '0')
}

func (l *strInt) ByteatIdx() byte {
	return l.Constant[l.Idx]
}

func (l *strInt) IntatIdx() int {
	if l.Len() == 0 || l.Idx < 0 {
		return 0
	}
	return int(l.ByteatIdx() - '0')
}

func (l *strInt) DeZeroes() {
	// 只剩一个0或找到数字即停
	for l.Len() != 1 && l.Intat(0) == 0 {
		l.Constant = l.Constant[1:]
	}
}

func (l *strInt) String() string {
	return l.Constant
}

func (l *strInt) deepCopy() *strInt {
	return &strInt{
		TempVal:      l.TempVal,
		Overflow: l.Overflow,
		Idx:      l.Idx,
		Constant: l.Constant,
	}
}

const digits = "0123456789"

func singleItoa(i int) string {
	return string(digits[i])
}
