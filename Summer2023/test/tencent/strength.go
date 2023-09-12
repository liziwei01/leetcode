/*
 * @Author: liziwei01
 * @Date: 2023-09-10 08:52:28
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-10 09:02:33
 * @Description: file content
 */
package tencent

import (
	"fmt"
	"sort"
)

func strength() {
	n := 0
	fmt.Scan(&n)
	powers := make([]int, 0)
	for i := 0; i != n; i++ {
		power := 0
		fmt.Scan(&power)
		powers = append(powers, power)
	}
	sort.Ints(powers)

	bulkStrength, bulkCourage := 0, 0
	for len(powers) > 1 {
		bulkStrength, bulkCourage = dealingwith(powers[len(powers)-1], bulkStrength, bulkCourage)
		bulkStrength, bulkCourage = dealingwith(powers[0], bulkStrength, bulkCourage)
	}
	if len(powers) == 1 {
		bulkStrength, bulkCourage = dealingwith(powers[0], bulkStrength, bulkCourage)
	}
	fmt.Printf("%d", bulkCourage)
}

func dealingwith(power, bulkStrength, bulkYq int) (int, int) {
	if bulkStrength <= power {
		bulkYq += abs(bulkStrength - power)
	}
	bulkStrength = power

	return bulkStrength, bulkYq
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}