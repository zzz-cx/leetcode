# 反转链表

> LeetCode 206 · [reverse-linked-list](https://leetcode.cn/problems/reverse-linked-list/)

## 题目

反转单链表，返回反转后的头节点。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
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
    def reverse_list(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head
        prev = None
        current = head
        while current is not None:
            nxt = current.next
            current.next = prev
            prev = current
            current = nxt
        return prev
```
