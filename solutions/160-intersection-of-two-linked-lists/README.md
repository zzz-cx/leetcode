# 相交链表

> LeetCode 160 · [intersection-of-two-linked-lists](https://leetcode.cn/problems/intersection-of-two-linked-lists/)

## 题目

给定两个单链表头节点，找出并返回两链表相交的起始节点；不相交则返回 null。

## 题解思路与解析

- 思路 哈希集合，如果节点不在哈希集合中，继续遍历到下一个节点；如果在，则后面的节点都在哈希集合中，从当前节点往后就都是相交节点
- 公共部分：8→4→5，两链表在此相交

## 解答

### Golang

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路 哈希集合，如果节点不在哈希集合中，继续遍历到下一个节点；如果在，则后面的节点都在哈希集合中，从当前节点往后就都是相交节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for node := headA; node != nil; node = node.Next {
		vis[node] = true
	}
	for node := headB; node != nil; node = node.Next {
		if vis[node] {
			return node
		}
	}
	return nil
}
```

### Python

```python
# 思路 哈希集合，如果节点不在哈希集合中，继续遍历到下一个节点；如果在，则后面的节点都在哈希集合中
from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def get_intersection_node(
        self, head_a: Optional[ListNode], head_b: Optional[ListNode]
    ) -> Optional[ListNode]:
        vis = set()
        node = head_a
        while node is not None:
            vis.add(node)
            node = node.next
        node = head_b
        while node is not None:
            if node in vis:
                return node
            node = node.next
        return None
```
