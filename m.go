/*
 * @Author: liziwei01
 * @Date: 2023-09-26 23:34:07
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-09-27 00:08:34
 * @Description: file content
 */
package leetcode

import (
	"fmt"
	"sort"
)

type elevator struct {
	status  uint  // 运行状态 0:停止关门;1:停止开门;2:上行;3:下行;4:检修
	current int   // 当前停留层数
	target  int   // 下一个目标层数
	targets queue // 所有目标层数
	
	// w // 荷载
}

type queue struct {
	q []int
	// q []struct
}

type Queue interface {
	pop() int
	top() int
	push(x int)
}

func (q *queue) pop() int {
	return 0
}

func (q *queue) top() int {
	return 0
}

func (q *queue) push(x int) {
}

type Elevator interface {
	StopAt(x int) error // 点击'x层'按钮，提示电梯需要在某层停下
	Open() error        // 开门按钮
	Close() error       // 关门按钮
	Call() error        // 物业沟通
	Targets() []int     // 电梯按钮常亮
	reach(x int) error  // 到达某层
	run(x uint) error   // 开始向下一个目标楼层行进

	// weight(x int) error // 
	// Emergency
	// Test
}

func (e *elevator) StopAt(x int) error {
	e.targets.push(x)
	// sort()
	return nil
}

func (e *elevator) Open() error {
	e.status = 1
	return nil
}

func (e *elevator) Close() error {
	e.status = 0
	return nil
}

func (e *elevator) Call() error {
	// 
	return nil
}

func (e *elevator) Targets() ([]int, error) {
	return e.targets.q, nil
}

func (e *elevator) reach(x int) error {
	e.current = x
	if e.status == 2 {
		// 上行则继续上行
		// 找到比当前层数大的所有层数中的最小值
		sort.Slice(e.targets, func(i, j int) bool { return e.targets.q[i] > e.targets.q[j] })
	} else {
		// 下行则继续下行
		// ...
		sort.Slice(e.targets.q, func(i, j int) bool { return e.targets.q[i] < e.targets.q[j] })
	}
	e.target = e.targets.pop()

	// 开门
	err := e.Open()
	if err != nil {
		// 报告物业等等流程
		return fmt.Errorf("某某原因")
	}
	e.Close()

	return e.run(e.target)
}

func (e *elevator) run(x int) error {
	if x < e.current {
		// 下行
		e.status = 3
	} else if x > e.current {
		// 上行
		e.status = 2
	} else {
		// 保持开门
		return e.Open()
	}

	return nil
}
