# 删除链表的倒数第N个结点

> LeetCode 19 · [remove-nth-node-from-end-of-list](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)

## 题目

给定链表头节点 `head` 和整数 `n`，删除链表的倒数第 `n` 个结点，并返回链表的头结点。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first := dummy
	second := dummy
	for i := 0; i < n+1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
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
    def remove_nth_from_end(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        first = dummy
        second = dummy
        for _ in range(n + 1):
            first = first.next
        while first is not None:
            first = first.next
            second = second.next
        second.next = second.next.next
        return dummy.next
```
