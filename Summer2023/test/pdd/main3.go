/*
 * @Author: liziwei01
 * @Date: 2023-08-22 08:41:50
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-08-22 08:56:10
 * @Description: file content
 */
package pdd

import "fmt"

func m(M int, Xs map[int]int) {
	// T := 0
	// fmt.Scan(&T)

	// for i := 0; i != T; i++ {
	// sum := 0
	// N, M := 0, 0
	// fmt.Scan(&N, &M)
	// Xs := make(map[int]int)
	// for j := 0; j != N; j++ {
	// 	X := 0
	// 	fmt.Scan(&X)
	// 	Xs[X]++
	// }

	// for range
	// }
}

func main3(M int, Xs map[int]int) {

	sum := 0
	N, M := 0, 0
	fmt.Scan(&N, &M)

	Msm := make(map[int]int)
	for j := 0; j != N; j++ {
		X := 0
		fmt.Scan(&X)

		if Msm[X] >= 0 {
			sum++
			Msm[X]--
			continue
		}

		if Xs[X] == 1 {
			sum++
			Xs[X]--
			continue
		}

		addMs(Msm, M, X)
		Xs[X]++
	}

	fmt.Printf("%d\n", sum)

}

func addMs(Ms map[int]int, M, X int) {
	originalM := M
	for M <= 100000 {
		if M-X >= 0 {
			Ms[M-X]++
		}
		M += originalM
	}
}
