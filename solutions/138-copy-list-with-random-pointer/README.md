# 复制带随机指针的链表

> LeetCode 138 · [copy-list-with-random-pointer](https://leetcode.cn/problems/copy-list-with-random-pointer/)

## 题目

深拷贝带随机指针的链表。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}
```

### Python

```python
from typing import Optional


class Node:
    def __init__(self, x: int, next=None, random=None):
        self.val = x
        self.next = next
        self.random = random


class Solution:
    def copy_random_list(self, head: Optional[Node]) -> Optional[Node]:
        if head is None:
            return None
        m = {}
        cur = head
        while cur is not None:
            m[cur] = Node(cur.val)
            cur = cur.next
        cur = head
        while cur is not None:
            m[cur].next = m.get(cur.next)
            m[cur].random = m.get(cur.random)
            cur = cur.next
        return m[head]
```
