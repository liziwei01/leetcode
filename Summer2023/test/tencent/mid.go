package tencent

import (
	"fmt"
	"sort"
)

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 0; i != t; i++ {
		var mid float64 = 0
		var midint int = 0

		n := 0
		fmt.Scan(&n)
		as := make([]int, 0)
		bs := make([]int, 0)
		for j := 0; j != n; j++ {
			tmp := 0
			fmt.Scan(&tmp)
			as = append(as, tmp)
		}
		for j := 0; j != n-1; j++ {
			tmp := 0
			fmt.Scan(&tmp)
			bs = append(bs, tmp)
		}
		sort.Ints(as)

		ntmp := n
		for j := 0; j != n-1; j++ {
			as = append(as[:bs[j]-j-1], as[bs[j]+1-j-1:]...)
			ntmp--
			if !isOdd(ntmp) {
				if (isOdd(as[ntmp/2-1]) == true && isOdd(as[ntmp/2]) == true) || (isOdd(as[ntmp/2-1]) == false && isOdd(as[ntmp/2]) == false) {
					midint = (as[ntmp/2-1] + as[ntmp/2]) / 2
					fmt.Printf("%d", midint)
				} else {
					mid = float64(as[ntmp/2-1]+as[ntmp/2]) / 2
					fmt.Printf("%.1f", mid)
				}
			} else {
				midint = as[ntmp/2]
				fmt.Printf("%d", midint)
			}
		}
		fmt.Printf("\n")
	}
}

func isOdd(i int) bool {
	return i%2 == 1
}
