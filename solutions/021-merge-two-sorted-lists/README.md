# 合并两个有序链表

> LeetCode 21 · [merge-two-sorted-lists](https://leetcode.cn/problems/merge-two-sorted-lists/)

## 题目

将两个升序链表合并为一个新的升序链表并返回。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
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
    def merge_two_lists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        tail = dummy
        while list1 is not None and list2 is not None:
            if list1.val < list2.val:
                tail.next = list1
                list1 = list1.next
            else:
                tail.next = list2
                list2 = list2.next
            tail = tail.next
        if list1 is not None:
            tail.next = list1
        else:
            tail.next = list2
        return dummy.next
```
