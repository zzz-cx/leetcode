# 环形链表 II

> LeetCode 142 · [linked-list-cycle-ii](https://leetcode.cn/problems/linked-list-cycle-ii/)

## 题目

给定链表头节点，返回链表开始入环的第一个节点；无环则返回 null。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
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
    def detect_cycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        fast, slow = head, head
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                p = head
                while p != slow:
                    p = p.next
                    slow = slow.next
                return p
        return None
```
