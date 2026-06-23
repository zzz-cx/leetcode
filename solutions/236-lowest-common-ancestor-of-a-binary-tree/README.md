# 二叉树的最近公共祖先

> LeetCode 236 · [lowest-common-ancestor-of-a-binary-tree](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/)

## 题目

给定二叉树根节点和两个节点 p、q，返回它们的最近公共祖先。

## 题解思路与解析

- 示例树（LeetCode 官方示例1）
- 3
- /   \
- 5     1
- / \   / \
- 6  2  0   8
- / \
- 7   4

## 解答

### Golang

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root // 空节点返回 nil；命中 p/q 则把“命中的节点”向上返回
	}
	left := lowestCommonAncestor(root.Left, p, q)   // 左子树里是否能找到 p 或 q（或它们的 LCA）
	right := lowestCommonAncestor(root.Right, p, q) // 右子树里是否能找到 p 或 q（或它们的 LCA）
	if left != nil && right != nil {
		return root // p、q 分居两侧（或一侧命中 p，另一侧命中 q），当前 root 即最近公共祖先
	}
	if left != nil {
		return left // 只在左边找到：把找到的节点/祖先继续向上冒泡
	}
	return right // 只在右边找到：把找到的节点/祖先继续向上冒泡
}
```

### Python

```python
from typing import Optional


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def lowest_common_ancestor(
        self, root: TreeNode, p: TreeNode, q: TreeNode
    ) -> Optional[TreeNode]:
        if root is None or root == p or root == q:
            return root
        left = self.lowest_common_ancestor(root.left, p, q)
        right = self.lowest_common_ancestor(root.right, p, q)
        if left is not None and right is not None:
            return root
        if left is not None:
            return left
        return right
```
