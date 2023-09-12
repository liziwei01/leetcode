package baidu

import (
	"fmt"
	"math/big"
)

func huati() {
	n := 0
	var tmp int64 = 0
	fmt.Scan(&n)
	as := make([]int64, 0)
	bs := make([]int64, 0)
	posDiff := big.NewInt(0)
	NegDiff := big.NewInt(0)
	
	for i := 0; i != n; i++ {
		fmt.Scan(&tmp)
		as = append(as, tmp)
	}
	for i := 0; i != n; i++ {
		fmt.Scan(&tmp)
		bs = append(bs, tmp)
		if as[i] - tmp > 0 {
			posDiff.Add(posDiff, big.NewInt(as[i] - tmp))
		} else {
			NegDiff.Add(NegDiff, big.NewInt(as[i] - tmp))
		}
	}
	
	NegDiff.Neg(NegDiff)

	if posDiff.Cmp(NegDiff) == 1 {
		fmt.Print(posDiff.Int64())
	} else {
		fmt.Print(NegDiff.Int64())
	}
}
