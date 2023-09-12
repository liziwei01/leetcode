package main

import (
	"fmt"
	"sort"
)

func main() {
	t := 1
	// fmt.Scan(&t)

	for i := 0; i != t; i++ {
		var mid float64 = 0
		var midint int = 0

		n := 5
		// fmt.Scan(&n)
		as := make([]int, 0)
		bs := make([]int, 0)
		as = append(as, 2)
		as = append(as, 2)
		as = append(as, 1)
		as = append(as, 3)
		as = append(as, 5)
		bs = append(bs, 3)
		bs = append(bs, 1)
		bs = append(bs, 2)
		bs = append(bs, 5)
		// for j := 0; j != n; j++ {
		// 	tmp := 0
		// 	fmt.Scan(&tmp)
		// 	as = append(as, tmp)
		// }
		// for j := 0; j != n-1; j++ {
		// 	tmp := 0
		// 	fmt.Scan(&tmp)
		// 	bs = append(bs, tmp)
		// }
		sort.Ints(as)

		ntmp := n
		if !isOdd(ntmp) {
			if (isOdd(as[ntmp/2-1]) == true && isOdd(as[ntmp/2]) == true) || (isOdd(as[ntmp/2-1]) == false && isOdd(as[ntmp/2]) == false) {
				midint = (as[ntmp/2-1] + as[ntmp/2]) / 2
				fmt.Printf("%d ", midint)
			} else {
				mid = float64(as[ntmp/2-1]+as[ntmp/2]) / 2
				fmt.Printf("%.1f ", mid)
			}
		} else {
			midint = as[ntmp/2]
			fmt.Printf("%d ", midint)
		}
		for j := 0; j != n-1; j++ {
			idx := bs[j]-1
			as = append(as[:idx], as[idx+1:]...)
			ntmp--
			if !isOdd(ntmp) {
				if (isOdd(as[ntmp/2-1]) == true && isOdd(as[ntmp/2]) == true) || (isOdd(as[ntmp/2-1]) == false && isOdd(as[ntmp/2]) == false) {
					midint = (as[ntmp/2-1] + as[ntmp/2]) / 2
					fmt.Printf("%d ", midint)
				} else {
					mid = float64(as[ntmp/2-1]+as[ntmp/2]) / 2
					fmt.Printf("%.1f ", mid)
				}
			} else {
				midint = as[ntmp/2]
				fmt.Printf("%d ", midint)
			}
		}
		fmt.Printf("\n")
	}
}

func isOdd(i int) bool {
	return i%2 == 1
}
