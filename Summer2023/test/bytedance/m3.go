package bytedance

import (
	"fmt"
	"math"
)

// 计算长度为 n 的正整数的权值之和
func calculateTotalWeight(n int) int {
	// 边界情况：n 小于等于 0 时，权值之和为 0
	if n <= 0 {
		return 0
	}

	// 计算 10^n 和 10^(n-1) 用于计算权值
	pow10n := int(math.Pow(10, float64(n)))
	pow10nMinus1 := int(math.Pow(10, float64(n-1)))

	// 奇数位和偶数位的权值贡献
	oddContribution := n * (pow10n + 1) / 2
	evenContribution := (n - 1) * pow10nMinus1 / 2

	// 计算权值之和
	totalWeight := oddContribution * evenContribution

	return totalWeight
}

func main33() {
	n := 3 // 输入 n 的值
	totalWeight := calculateTotalWeight(n)
	fmt.Printf("Total weight sum for %d-digit numbers: %d\n", n, totalWeight)
}

