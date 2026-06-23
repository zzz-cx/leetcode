# 路径总和 III

> LeetCode 437 · [path-sum-iii](https://leetcode.cn/problems/path-sum-iii/)

## 题目

给定二叉树和整数 `targetSum`，返回路径和等于 `targetSum` 的路径数目（路径从根到叶）。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	Val := root.Val
	if Val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-Val)
	res += rootSum(root.Right, targetSum-Val)
	return
}
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
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
    def root_sum(self, root: Optional[TreeNode], target_sum: int) -> int:
        if root is None:
            return 0
        val = root.val
        res = 0
        if val == target_sum:
            res += 1
        res += self.root_sum(root.left, target_sum - val)
        res += self.root_sum(root.right, target_sum - val)
        return res

    def path_sum(self, root: Optional[TreeNode], target_sum: int) -> int:
        if root is None:
            return 0
        res = self.root_sum(root, target_sum)
        res += self.path_sum(root.left, target_sum)
        res += self.path_sum(root.right, target_sum)
        return res
```
