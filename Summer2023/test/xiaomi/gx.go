/*
 * @Author: liziwei01
 * @Date: 2023-09-23 05:07:16
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-23 05:23:20
 * @Description: file content
 */
package xiaomi

import (
	"fmt"
	"strings"
)

func main() {
	n := 0
	fmt.Scanln(&n)
	seqIn := ""
	fmt.Scanln(&seqIn)
	seqs := strings.Split(seqIn, ",")
	m := make(map[string][]string)
	for i := 0; i != len(seqs); i++ {
		prereq := strings.Split(seqs[i], ":")
		pre := prereq[0]
		req := prereq[1]
		m[pre] = append(m[pre], req)
	}

	canF := true
	visited := make(map[string]bool)
	var visit func(node string)
	visit = func(node string) {
		if visited[node] {
			canF = false
			return
		}
		visited[node] = true
		for _, v := range m[node] {
			visit(v)
			visited[node] = false
		}
	}
	for k, _ := range m {
		visit(k)
	}
	if canF {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}
