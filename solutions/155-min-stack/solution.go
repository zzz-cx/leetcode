package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.minStack) > 0 && top == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	min1 := s.GetMin()
	s.Pop()
	top := s.Top()
	min2 := s.GetMin()
	status := "PASS"
	if min1 != -3 || top != 0 || min2 != -2 {
		status = "FAIL"
	}
	fmt.Printf("%s | GetMin=%d Top=%d GetMin=%d\n", status, min1, top, min2)
}
