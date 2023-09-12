package baidu

import (
	"fmt"
)

func main() {
	n, m, l, r := 0, 0, 0, 0
	fmt.Scan(&n, &m)

	var a int64 = 0
	as := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		as[i] = a
	}

	for i := 1; i <= m; i++ {
		fmt.Scan(&l, &r)

		for j := l; j <= r; j++ {
			as[j]--
			if as[j] < 0 {
				fmt.Print(i - 1)
				return
			}
		}
	}
	fmt.Print(m)
}
