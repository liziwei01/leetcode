/*
 * @Author: liziwei01
 * @Date: 2023-09-23 07:15:09
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-23 08:04:15
 * @Description: file content
 */
package oppo

import (
	"fmt"
	"math/big"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)
	as := make([]int, n)
	bs := make([]int, n)
	// as[0] = 1
	// as[1] = 2
	// as[2] = 3
	// as[3] = 4
	// bs[0] = 1
	// bs[1] = 3
	// bs[2] = 2
	// bs[3] = 1

	var suma, sumb, mina, maxa, minb, maxb, minaidx, minbidx, maxaidx, maxbidx int = 0, 0, 100000000000000, 0, 100000000000000, 0, 0, 0, 0, 0
	for i := 0; i != n; i++ {
		fmt.Scan(&as[i], &bs[i])
		suma += as[i]
		sumb += bs[i]
		if mina > as[i] {
			mina = as[i]
			minaidx = i
		}
		if minb > bs[i] {
			minb = bs[i]
			minbidx = i
		}
		if maxa < as[i] {
			maxa = as[i]
			maxaidx = i
		}
		if maxb < bs[i] {
			maxb = bs[i]
			maxbidx = i
		}
	}

	sortedas := deepCopy(as)
	sortedbs := deepCopy(bs)
	sort.Ints(sortedas)
	sort.Ints(sortedbs)

	for i := 0; i != n; i++ {
		tmpSuma := big.NewFloat(float64(suma))
		tmpSumb := big.NewFloat(float64(sumb))
		// ta, _ := tmpSuma.Float64()
		// tb, _ := tmpSumb.Float64()
		tmpSuma.Sub(tmpSuma, big.NewFloat(float64(as[i])))
		tmpSumb.Sub(tmpSumb, big.NewFloat(float64(bs[i])))
		// ta, _ = tmpSuma.Float64()
		// tb, _ = tmpSumb.Float64()
		if i != minaidx {
			tmpSuma.Sub(tmpSuma, big.NewFloat(float64(mina)))
		} else {
			// 最低外观分(mina)评委失去资格
			tmpSuma.Sub(tmpSuma, big.NewFloat(float64(sortedas[1])))
		}
		// ta, _ = tmpSuma.Float64()
		// tb, _ = tmpSumb.Float64()
		if i != maxaidx {
			tmpSuma.Sub(tmpSuma, big.NewFloat(float64(maxa)))
		} else {
			// 最高外观分(maxa)评委失去资格
			tmpSuma.Sub(tmpSuma, big.NewFloat(float64(sortedas[n-2])))
		}
		// ta, _ = tmpSuma.Float64()
		// tb, _ = tmpSumb.Float64()
		if i != minbidx {
			tmpSumb.Sub(tmpSumb, big.NewFloat(float64(minb)))
		} else {
			// 最低性能分(minb)评委失去资格
			tmpSumb.Sub(tmpSumb, big.NewFloat(float64(sortedbs[1])))
		}
		// ta, _ = tmpSuma.Float64()
		// tb, _ = tmpSumb.Float64()
		if i != maxbidx {
			tmpSumb.Sub(tmpSumb, big.NewFloat(float64(maxb)))
		} else {
			// 最高性能分(maxb)评委失去资格
			tmpSumb.Sub(tmpSumb, big.NewFloat(float64(sortedbs[n-2])))
		}
		// ta, _ = tmpSuma.Float64()
		// tb, _ = tmpSumb.Float64()

		tmpSuma.Quo(tmpSuma, big.NewFloat(float64(n-1-2)))
		tmpSumb.Quo(tmpSumb, big.NewFloat(float64(n-1-2)))
		// ta, _ = tmpSuma.Float64()
		// tb, _ = tmpSumb.Float64()
		// fmt.Println(ta)
		// fmt.Println(tb)
		sum := big.NewFloat(0).Add(tmpSuma, tmpSumb)
		sum.Quo(sum, big.NewFloat(2))
		sumf, _ := sum.Float64()
		fmt.Println(sumf)
	}
}

func deepCopy(s []int) []int {
	newS := make([]int, len(s))
	copy(newS, s)
	return newS
}
