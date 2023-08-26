package pdd

import (
	"math/big"
)

func main2(N int) int64 {
	// N := 0
	// fmt.Scan(&N)
	var sum = big.NewInt(0)

	// i = red number
	redMax := N / 2
	if isOdd(N) {
		redMax++
	}

	for i := 0; i <= redMax; i++ {
		red := i
		white := N - i
		whiteInterval := white + 1

		bi := comb(whiteInterval, red)
		sum.Add(sum, bi)
	}

	bigModules := big.NewInt(int64(1000000000+7))
	sum.Mod(sum, bigModules)

	// fmt.Print(sum)

	return sum.Int64()
}

func isOdd(i int) bool {
	return i%2 == 1
}

// n 选 k
func comb(n, k int) *big.Int {
	return new(big.Int).Binomial(int64(n), int64(k))
}
