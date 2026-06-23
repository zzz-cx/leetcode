# 最小栈

> LeetCode 155 · [min-stack](https://leetcode.cn/problems/min-stack/)

## 题目

设计支持 push、pop、top 操作，并在常数时间内检索最小元素的栈。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
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

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
```

### Python

```python
class MinStack:
    def __init__(self):
        self.stack = []
        self.min_stack = []

    def push(self, val: int) -> None:
        self.stack.append(val)
        if not self.min_stack or val <= self.min_stack[-1]:
            self.min_stack.append(val)

    def pop(self) -> None:
        if not self.stack:
            return
        top = self.stack.pop()
        if self.min_stack and top == self.min_stack[-1]:
            self.min_stack.pop()

    def top(self) -> int:
        return self.stack[-1]

    def get_min(self) -> int:
        return self.min_stack[-1]
```
