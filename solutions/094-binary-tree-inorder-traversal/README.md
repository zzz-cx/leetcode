# 二叉树的中序遍历

> LeetCode 94 · [binary-tree-inorder-traversal](https://leetcode.cn/problems/binary-tree-inorder-traversal/)

## 题目

给定二叉树根节点，返回其中序遍历结果。

## 题解思路与解析

- 思路：递归，先遍历左子树，再遍历根节点，再遍历右子树
- 这里使用外部out＋内部递归函数的形式

## 解答

### Golang

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：递归，先遍历左子树，再遍历根节点，再遍历右子树
//这里使用外部out＋内部递归函数的形式
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := inorderTraversal(root.Left)
	out = append(out, root.Val)
	right := inorderTraversal(root.Right)
	out = append(out, right...)
	return out
}
```

### Python

```python
# 思路：递归，先遍历左子树，再遍历根节点，再遍历右子树
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def inorder_traversal(self, root: Optional[TreeNode]) -> List[int]:
        if root is None:
            return []
        out = self.inorder_traversal(root.left)
        out.append(root.val)
        right = self.inorder_traversal(root.right)
        out.extend(right)
        return out
```
