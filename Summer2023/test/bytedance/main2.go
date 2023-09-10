package bytedance

import (
    "fmt"
)

func main2() {
    n := 0
    m := 0
    fmt.Scan(&n, &m)
    matrix := make([][]int, 0)
    for i := 0; i != n; i++ {
        mStrings := make([]int, 0)
        for j := 0; j != m; j++ {
            jString := ""
            fmt.Scan(&jString)
            if jString == "red" {
                mStrings = append(mStrings, 0)
            }
            if jString == "green" {
                mStrings = append(mStrings, 1)
            }
            if jString == "blue" {
                mStrings = append(mStrings, 2)
            }
        }
        matrix = append(matrix, mStrings)
    }
    q := 0
    fmt.Scan(&q)
    
    for i := 0; i != q; i++ {
        x1, y1, x2, y2 := 0, 0, 0, 0
        fmt.Scan(&x1, &y1, &x2, &y2)
        n := colorNum(matrix, x1-1, y1-1, x2-1, y2-1)
        fmt.Printf("%d\n", n)
    }

}

func colorNum(matrix [][]int, x1, y1, x2, y2 int) int {
    r, g, b := 0, 0, 0
    for i := x1; i <= x2; i++ {
        for j := y1; j <= y2; j++ {
            if matrix[i][j] == 0 {
                r = 1
            } else if matrix[i][j] == 1 {
                g = 1
            } else if matrix[i][j] == 2 {
                b = 1
            }
        }
    }
    return r + g + b
}
