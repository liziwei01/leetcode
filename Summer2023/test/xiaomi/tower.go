/*
 * @Author: liziwei01
 * @Date: 2023-09-13 16:56:04
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-23 05:07:00
 * @Description: file content
 */
package xiaomi

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type tower struct {
	x int
	y int
	q int
}

func to() {
	towerLength, radius := 0, 0
	line := ""
	fmt.Scanln(&line)
	sli := strings.Split(line, ",")
	towerLength, _ = strconv.Atoi(sli[0])
	radius, _ = strconv.Atoi(sli[1])
	// fmt.Printf("%d,%d\n", towerLength, radius)
	xmax, ymax := 0, 0
	towers := make([]tower, towerLength)
	for i := 0; i != towerLength; i++ {
		x, y, q := 0, 0, 0
		fmt.Scanln(&line)
		sli = strings.Split(line, ",")
		x, _ = strconv.Atoi(sli[0])
		y, _ = strconv.Atoi(sli[1])
		q, _ = strconv.Atoi(sli[2])
		// fmt.Printf("%d,%d,%d\n", x, y, q)
		if x > xmax {
			xmax = x
		}
		if y > ymax {
			ymax = y
		}
		towers[i] = tower{
			x: x,
			y: y,
			q: q,
		}
	}
	// towers[0] = tower{1, 2, 5}
	// towers[1] = tower{2, 1, 7}
	// towers[2] = tower{3, 1, 9}
	signal := make([][]int64, xmax)
	var maxSig int64 = 0
	maxSigxs := make([]int, 0)
	maxSigys := make([]int, 0)
	for i := 0; i != xmax; i++ {
		signal[i] = make([]int64, ymax)
		for j := 0; j != ymax; j++ {
			var sig int64 = 0
			for k := 0; k != towerLength; k++ {
				dis := distance(i, j, towers[k])
				sig += sign(dis, towers[k], radius)
			}
			signal[i][j] = sig
			if sig > maxSig {
				maxSig = sig
				maxSigxs = make([]int, 0)
				maxSigys = make([]int, 0)
				maxSigxs = append(maxSigxs, i)
				maxSigys = append(maxSigys, j)
			}
			if sig == maxSig {
				maxSigxs = append(maxSigxs, i)
				maxSigys = append(maxSigys, j)
			}
		}
	}
	minx, miny, minSum := xmax, ymax, 1000000000
	for i := 0; i != len(maxSigxs); i++ {
		s := maxSigxs[i] + maxSigys[i]
		if s < minSum {
			minSum = s
			minx = maxSigxs[i]
			miny = maxSigys[i]
		}
	}
	fmt.Printf("%d,%d", minx, miny)
}

func distance(x, y int, t tower) *big.Float {
	dx := abs(x - t.x)
	dy := abs(y - t.y)
	d := big.NewFloat(float64(dx*dx + dy*dy))
	return d.Sqrt(d)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sign(d *big.Float, t tower, r int) int64 {
	if d.Cmp(big.NewFloat(float64(r))) == 1 {
		// 大于
		return 0
	}
	s := new(big.Float).Quo(big.NewFloat(float64(t.q)), d.Add(d, big.NewFloat(1)))
	sFloor, _ := s.Int64()
	return sFloor
}
