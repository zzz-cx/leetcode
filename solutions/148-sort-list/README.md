# 排序链表

> LeetCode 148 · [sort-list](https://leetcode.cn/problems/sort-list/)

## 题目

对链表进行升序排序，要求 O(n log n) 时间、O(1) 空间（常数级额外空间）。

## 题解思路与解析

- sortList 排序链表（LeetCode 148）
- 思路：归并排序。找中点拆成左右两段 → 递归排序 → mergeTwoLists 合并
- 时间 O(n log n)，空间 O(log n) 递归栈（满足题面 O(1) 空间需自底向上归并，面试写这版即可）
- getMid 快慢指针找中点，并在 slow 处断开链表
- slow 最终停在左半最后一个节点；mid = slow.Next 为右半头；slow.Next = nil 切断
- func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
- dummy := &ListNode{0, nil}
- tail := dummy
- for list1 != nil && list2 != nil {
- if list1.Val < list2.Val {
- tail.Next = list1
- list1 = list1.Next
- } else {
- tail.Next = list2
- list2 = list2.Next
- }
- tail = tail.Next
- if list1 != nil {
- return dummy.Next

## 解答

### Golang

```go
// sortList 排序链表（LeetCode 148）
// 思路：归并排序。找中点拆成左右两段 → 递归排序 → mergeTwoLists 合并
// 时间 O(n log n)，空间 O(log n) 递归栈（满足题面 O(1) 空间需自底向上归并，面试写这版即可）
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head // 0 个或 1 个节点，已有序
	}
	mid := getMid(head)               // 从 mid 处断开，左半 head..slow，右半 mid..末尾
	left := sortList(head)            // 排左半
	right := sortList(mid)            // 排右半
	return mergeTwoLists(left, right) // 合并两个有序链表（见 21mergeTwoLists.go）
}

// getMid 快慢指针找中点，并在 slow 处断开链表
// slow 最终停在左半最后一个节点；mid = slow.Next 为右半头；slow.Next = nil 切断
func getMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next // fast 比 slow 多走一步，保证左半不少于右半
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mid
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	dummy := &ListNode{0, nil}
// 	tail := dummy
// 	for list1 != nil && list2 != nil {
// 		if list1.Val < list2.Val {
// 			tail.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			tail.Next = list2
// 			list2 = list2.Next
// 		}
// 		tail = tail.Next
// 	}
// 	if list1 != nil {
// 		tail.Next = list1
// 	} else {
// 		tail.Next = list2
// 	}
// 	return dummy.Next
// }
```

### Python

```python
# sort_list 排序链表（LeetCode 148）
# 思路：归并排序。找中点拆成左右两段 → 递归排序 → merge_two_lists 合并
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def sort_list(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head
        mid = self._get_mid(head)
        left = self.sort_list(head)
        right = self.sort_list(mid)
        return self._merge_two_lists(left, right)

    def _get_mid(self, head: ListNode) -> Optional[ListNode]:
        slow, fast = head, head.next
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
        mid = slow.next
        slow.next = None
        return mid

    def _merge_two_lists(
        self, list1: Optional[ListNode], list2: Optional[ListNode]
    ) -> Optional[ListNode]:
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
