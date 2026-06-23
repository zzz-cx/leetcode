# 两数相加

> LeetCode 2 · [add-two-numbers](https://leetcode.cn/problems/add-two-numbers/)

## 题目

给你两个非空的链表，表示两个非负的整数，数字按逆序存储，每个节点存储一个数字。将两个数相加并以相同形式返回一个表示和的链表。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cnt := 0
	dummy := new(ListNode)
	tail := dummy
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + cnt
		cnt = sum / 10
		sum = sum % 10
		tail.Next = &ListNode{Val: sum}
		tail = tail.Next
	}
	if cnt > 0 {
		tail.Next = &ListNode{Val: cnt}
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
    def add_two_numbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        cnt = 0
        dummy = ListNode()
        tail = dummy
        while l1 is not None or l2 is not None:
            n1, n2 = 0, 0
            if l1 is not None:
                n1 = l1.val
                l1 = l1.next
            if l2 is not None:
                n2 = l2.val
                l2 = l2.next
            s = n1 + n2 + cnt
            cnt = s // 10
            s = s % 10
            tail.next = ListNode(s)
            tail = tail.next
        if cnt > 0:
            tail.next = ListNode(cnt)
        return dummy.next
```
