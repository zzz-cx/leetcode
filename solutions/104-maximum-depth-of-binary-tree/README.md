# 二叉树的最大深度

> LeetCode 104 · [maximum-depth-of-binary-tree](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)

## 题目

给定二叉树根节点，返回其最大深度。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
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
    def max_depth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        left_depth = self.max_depth(root.left)
        right_depth = self.max_depth(root.right)
        return max(left_depth, right_depth) + 1
```
