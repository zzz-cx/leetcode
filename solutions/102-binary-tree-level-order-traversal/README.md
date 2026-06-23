# 二叉树的层序遍历

> LeetCode 102 · [binary-tree-level-order-traversal](https://leetcode.cn/problems/binary-tree-level-order-traversal/)

## 题目

给定二叉树根节点，返回其层序遍历结果（逐层从左到右）。

## 题解思路与解析

- 递归解法
- 广度优先遍历（用队列）

## 解答

### Golang

```go
// 递归解法
func helper(node *TreeNode, level int, ret [][]int) [][]int {
	if node == nil {
		return ret
	}
	if len(ret) == level {
		ret = append(ret, []int{}) // 如果当前层数没有对应的切片，则创建一个
	}
	ret[level] = append(ret[level], node.Val) // 将当前节点的值添加到当前层对应的切片中
	ret = helper(node.Left, level+1, ret)
	ret = helper(node.Right, level+1, ret)
	return ret
}
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	return helper(root, 0, ret)
}

// 广度优先遍历（用队列）
```

### Python

```python
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def _helper(self, node: Optional[TreeNode], level: int, ret: List[List[int]]) -> List[List[int]]:
        if node is None:
            return ret
        if len(ret) == level:
            ret.append([])
        ret[level].append(node.val)
        ret = self._helper(node.left, level + 1, ret)
        ret = self._helper(node.right, level + 1, ret)
        return ret

    def level_order(self, root: Optional[TreeNode]) -> List[List[int]]:
        ret: List[List[int]] = []
        if root is None:
            return ret
        return self._helper(root, 0, ret)
```
