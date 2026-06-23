# 第K小的元素

> LeetCode 230 · [kth-smallest-element-in-a-bst](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/)

## 题目

给定 BST 根节点和 k，返回其中第 k 小的元素。

## 题解思路与解析

- 思路：中序遍历，记录遍历的节点数量，当遍历到第k个节点时，返回该节点的值
- 构造测试用例
- 调用 kthSmallest 并打印结果

## 解答

### Golang

```go
// 思路：中序遍历，记录遍历的节点数量，当遍历到第k个节点时，返回该节点的值
func kthSmallest(root *TreeNode, k int) int {
	count := 0
	var result int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return result
}
```

### Python

```python
# 思路：中序遍历，记录遍历的节点数量，当遍历到第k个节点时，返回该节点的值
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def kth_smallest(self, root: Optional[TreeNode], k: int) -> int:
        count = 0
        result = 0

        def dfs(node: Optional[TreeNode]) -> None:
            nonlocal count, result
            if node is None:
                return
            dfs(node.left)
            count += 1
            if count == k:
                result = node.val
                return
            dfs(node.right)

        dfs(root)
        return result
```
