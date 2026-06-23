# 二叉树展开为链表

> LeetCode 114 · [flatten-binary-tree-to-linked-list](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/)

## 题目

将二叉树展开为单链表，展开后的链表节点顺序与前序遍历相同。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = left
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
```

### Python

```python
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        if root is None:
            return
        self.flatten(root.left)
        self.flatten(root.right)
        left = root.left
        right = root.right
        root.left = None
        root.right = left
        p = root
        while p.right is not None:
            p = p.right
        p.right = right
```
