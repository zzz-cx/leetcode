# 两两交换链表中的节点

> LeetCode 24 · [swap-nodes-in-pairs](https://leetcode.cn/problems/swap-nodes-in-pairs/)

## 题目

两两交换相邻节点，返回交换后链表的头节点。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	current := head
	for current != nil && current.Next != nil {
		next := current.Next
		current.Next = next.Next
		next.Next = current
		prev.Next = next
		prev = current
		current = current.Next
	}
	return dummy.Next
}
```

### Python

```python
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def swap_pairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        prev = dummy
        current = head
        while current is not None and current.next is not None:
            nxt = current.next
            current.next = nxt.next
            nxt.next = current
            prev.next = nxt
            prev = current
            current = current.next
        return dummy.next
```
