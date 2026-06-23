# 翻转二叉树

> LeetCode 226 · [invert-binary-tree](https://leetcode.cn/problems/invert-binary-tree/)

## 题目

翻转二叉树，返回根节点。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root

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
    def invert_tree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if root is None:
            return None
        left = self.invert_tree(root.left)
        right = self.invert_tree(root.right)
        root.left = right
        root.right = left
        return root
```
