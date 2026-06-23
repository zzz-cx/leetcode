# 环形链表

> LeetCode 141 · [linked-list-cycle](https://leetcode.cn/problems/linked-list-cycle/)

## 题目

判断链表是否有环。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```

### Python

```python
from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def has_cycle(self, head: Optional[ListNode]) -> bool:
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                return True
        return False
```
