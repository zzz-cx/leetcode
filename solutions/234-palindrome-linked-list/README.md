# 回文链表

> LeetCode 234 · [palindrome-linked-list](https://leetcode.cn/problems/palindrome-linked-list/)

## 题目

判断链表是否为回文链表。

## 题解思路与解析

- isPalindrome 简单版（推荐面试先写）：链表值拷到切片，左右指针比较
- 时间 O(n)，空间 O(n)。代码短、不易错。
- isPalindromeO1 进阶版：快慢指针找中点 + 反转后半 + 比较，空间 O(1)

## 解答

### Golang

```go
// isPalindrome 简单版（推荐面试先写）：链表值拷到切片，左右指针比较
// 时间 O(n)，空间 O(n)。代码短、不易错。
func isPalindrome(head *ListNode) bool {
	vals := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		if vals[i] != vals[j] {
			return false
		}
	}
	return true
}

// isPalindromeO1 进阶版：快慢指针找中点 + 反转后半 + 比较，空间 O(1)
func isPalindromeO1(head *ListNode) bool {
	if head == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)
	p1, p2 := head, secondHalfStart
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1, p2 = p1.Next, p2.Next
	}
	firstHalfEnd.Next = reverseList(secondHalfStart) // 恢复链表（面试可省略）
	return true
}

func endOfFirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
```

### Python

```python
# is_palindrome 简单版（推荐面试先写）：链表值拷到切片，左右指针比较
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def is_palindrome(self, head: Optional[ListNode]) -> bool:
        vals = []
        cur = head
        while cur is not None:
            vals.append(cur.val)
            cur = cur.next
        i, j = 0, len(vals) - 1
        while i < j:
            if vals[i] != vals[j]:
                return False
            i += 1
            j -= 1
        return True

    def is_palindrome_o1(self, head: Optional[ListNode]) -> bool:
        if head is None:
            return True
        first_half_end = self._end_of_first_half(head)
        second_half_start = self._reverse_list(first_half_end.next)
        p1, p2 = head, second_half_start
        while p2 is not None:
            if p1.val != p2.val:
                return False
            p1, p2 = p1.next, p2.next
        first_half_end.next = self._reverse_list(second_half_start)
        return True

    def _end_of_first_half(self, head: ListNode) -> ListNode:
        slow, fast = head, head
        while fast.next is not None and fast.next.next is not None:
            slow = slow.next
            fast = fast.next.next
        return slow

    def _reverse_list(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev = None
        current = head
        while current is not None:
            nxt = current.next
            current.next = prev
            prev = current
            current = nxt
        return prev
```
