/*
 * @Author: liziwei01
 * @Date: 2023-09-23 02:19:11
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-23 02:24:40
 * @Description: file content
 */
package netease

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

func similar() {
	n := 0
	fmt.Scan(&n)
	sim := make(map[string]int)
	for i := 0; i != n; i++ {
		s := ""
		fmt.Scan(&s)
		divStr := strings.Split(s, "")
		sort.Strings(divStr)
		s = strings.Join(divStr, "")
		sim[s]++
	}

	sum := big.NewInt(0)
	for _, v := range sim {
		sum.Add(sum, comb(v, 2))
	}
	fmt.Print(sum.Int64())
}

// n 选 k
func comb(n, k int) *big.Int {
	return new(big.Int).Binomial(int64(n), int64(k))
}
