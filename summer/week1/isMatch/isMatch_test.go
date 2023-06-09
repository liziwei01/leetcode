/*
 * @Author: liziwei01
 * @Date: 2023-06-09 19:00:13
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-09 19:07:42
 * @Description: file content
 */
package ismatch

import (
	"testing"
	"time"
)

func TestIsMatch(t *testing.T) {
	start := time.Now()
	b := isMatch("babbbbaabababaabbababaababaabbaabababbaaababbababaaaaaabbabaaaabababbabbababbbaaaababbbabbbbbbbbbbaabbb", 
	"b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a")
	delta := time.Now().Sub(start)
	t.Log(b)
	t.Log(delta)
}