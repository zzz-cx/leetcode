# 二叉树的右视图

> LeetCode 199 · [binary-tree-right-side-view](https://leetcode.cn/problems/binary-tree-right-side-view/)

## 题目

给定二叉树根节点，返回从右侧能看到的节点值（层序最右）。

## 题解思路与解析

- 广度优先遍历，保留每层最后一个元素

## 解答

### Golang

```go
//广度优先遍历，保留每层最后一个元素
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	result := []int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == levelSize-1 {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}
```

### Python

```python
# 广度优先遍历，保留每层最后一个元素
from collections import deque
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def right_side_view(self, root: Optional[TreeNode]) -> List[int]:
        if root is None:
            return []
        queue = deque([root])
        result = []
        while queue:
            level_size = len(queue)
            for i in range(level_size):
                node = queue.popleft()
                if i == level_size - 1:
                    result.append(node.val)
                if node.left is not None:
                    queue.append(node.left)
                if node.right is not None:
                    queue.append(node.right)
        return result
```
