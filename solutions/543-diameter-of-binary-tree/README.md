# 二叉树的直径

> LeetCode 543 · [diameter-of-binary-tree](https://leetcode.cn/problems/diameter-of-binary-tree/)

## 题目

给定二叉树根节点，返回树的直径（任意两节点间最长路径的边数）。

## 题解思路与解析

- 思路：递归，计算左右子树的深度，然后计算直径
- 直径不一定经过根节点，因此需要遍历整棵树，在每个节点处用
- “左子树深度 + 右子树深度” 更新全局最大值。

## 解答

### Golang

```go
// 思路：递归，计算左右子树的深度，然后计算直径
func diameterOfBinaryTree(root *TreeNode) int {
	// 直径不一定经过根节点，因此需要遍历整棵树，在每个节点处用
	// “左子树深度 + 右子树深度” 更新全局最大值。
	diameter := 0

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		if left+right > diameter {
			diameter = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}

	depth(root)
	return diameter
}
```

### Python

```python
# 思路：递归，计算左右子树的深度，然后计算直径
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def diameter_of_binary_tree(self, root: Optional[TreeNode]) -> int:
        diameter = 0

        def depth(node: Optional[TreeNode]) -> int:
            nonlocal diameter
            if node is None:
                return 0
            left = depth(node.left)
            right = depth(node.right)
            if left + right > diameter:
                diameter = left + right
            return max(left, right) + 1

        depth(root)
        return diameter
```
