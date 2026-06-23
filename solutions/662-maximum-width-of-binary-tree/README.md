# 二叉树最大宽度

> LeetCode 662 · [maximum-width-of-binary-tree](https://leetcode.cn/problems/maximum-width-of-binary-tree/)

## 题目

给定二叉树根节点，返回其最大宽度（层中最左与最右节点下标差 + 1）。

## 题解思路与解析

- Pair 队列元素：节点 + 它在「满二叉树编号」下的下标
- 根编号 1，左子 2*idx，右子 2*idx+1（与层序下标一致，便于算层宽）
- widthOfBinaryTree 二叉树最大宽度（LeetCode 662）
- 思路：BFS 按层遍历。每一层用队首、队尾节点的 idx 算宽度 = right - left + 1，再取全局最大。
- 层宽定义：该层最左、最右「仍存在节点」之间的间隔（含中间空格，按满二叉树位置算）。
- 本层第一个、最后一个在队列里的下标 → 即该层最左、最右节点的编号
- 弹出本层所有节点，把下一层子节点入队
- 层序：1 / 2,3 / 4,5,6,7 → 最宽一层为第 3 层，宽度 4

## 解答

### Golang

```go
// Pair 队列元素：节点 + 它在「满二叉树编号」下的下标
// 根编号 1，左子 2*idx，右子 2*idx+1（与层序下标一致，便于算层宽）
type Pair struct {
	node *TreeNode
	idx  int
}

// widthOfBinaryTree 二叉树最大宽度（LeetCode 662）
//
// 思路：BFS 按层遍历。每一层用队首、队尾节点的 idx 算宽度 = right - left + 1，再取全局最大。
// 层宽定义：该层最左、最右「仍存在节点」之间的间隔（含中间空格，按满二叉树位置算）。
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []Pair{{root, 1}} // 根下标从 1 开始
	ans := 0

	for len(queue) > 0 {
		size := len(queue) // 当前层节点数，固定本轮只处理这么多

		// 本层第一个、最后一个在队列里的下标 → 即该层最左、最右节点的编号
		left := queue[0].idx
		right := queue[size-1].idx
		ans = max(ans, right-left+1)

		// 弹出本层所有节点，把下一层子节点入队
		for i := 0; i < size; i++ {
			pair := queue[0]
			queue = queue[1:]

			node, idx := pair.node, pair.idx
			if node.Left != nil {
				queue = append(queue, Pair{node.Left, idx * 2})
			}
			if node.Right != nil {
				queue = append(queue, Pair{node.Right, idx*2 + 1})
			}
		}
	}

	return ans
}
```

### Python

```python
# width_of_binary_tree 二叉树最大宽度（LeetCode 662）
# 思路：BFS 按层遍历。每一层用队首、队尾节点的 idx 算宽度 = right - left + 1，再取全局最大。
from collections import deque
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def width_of_binary_tree(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0

        queue = deque([(root, 1)])
        ans = 0

        while queue:
            size = len(queue)
            left = queue[0][1]
            right = queue[size - 1][1]
            ans = max(ans, right - left + 1)

            for _ in range(size):
                node, idx = queue.popleft()
                if node.left is not None:
                    queue.append((node.left, idx * 2))
                if node.right is not None:
                    queue.append((node.right, idx * 2 + 1))

        return ans
```
