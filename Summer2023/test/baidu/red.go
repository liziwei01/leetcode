package baidu

import (
    "fmt"
)

func red() {
    t := 0
    fmt.Scan(&t)

	for i := 0; i != t; i++ {
		n, m := 0, 0
		fmt.Scan(&n, &m)
		// 棋盘只有一边是偶数就能赢
		if (isEven(n) == true && isEven(m) == false) || (isEven(m) == true && isEven(n) == false) {
			fmt.Print("Yes\n")
		} else {
			fmt.Print("No\n")
		}
	}
}

func isEven(i int) bool {
	return i % 2 == 0
}