# 从前序与中序遍历序列构造二叉树

> LeetCode 105 · [construct-binary-tree-from-preorder-and-inorder-traversal](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

## 题目

根据前序遍历和中序遍历构造二叉树。

## 题解思路与解析

- buildTree2 使用哈希表 + 下标边界，整体时间复杂度 O(n)。

## 解答

### Golang

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			rootIndex = i
		}
	}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

// buildTree2 使用哈希表 + 下标边界，整体时间复杂度 O(n)。
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	pos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		pos[v] = i
	}

	preIdx := 0
	var helper func(l, r int) *TreeNode
	helper = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		val := preorder[preIdx]
		preIdx++

		root := &TreeNode{Val: val}
		mid := pos[val]
		root.Left = helper(l, mid-1)
		root.Right = helper(mid+1, r)
		return root
	}

	return helper(0, len(inorder)-1)
}
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
    def build_tree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0 or len(inorder) == 0:
            return None
        root = TreeNode(preorder[0])
        root_index = 0
        for i in range(len(inorder)):
            if inorder[i] == root.val:
                root_index = i
        root.left = self.build_tree(preorder[1 : root_index + 1], inorder[:root_index])
        root.right = self.build_tree(preorder[root_index + 1 :], inorder[root_index + 1 :])
        return root

    def build_tree2(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0:
            return None

        pos = {v: i for i, v in enumerate(inorder)}
        pre_idx = 0

        def helper(l: int, r: int) -> Optional[TreeNode]:
            nonlocal pre_idx
            if l > r:
                return None
            val = preorder[pre_idx]
            pre_idx += 1
            root = TreeNode(val)
            mid = pos[val]
            root.left = helper(l, mid - 1)
            root.right = helper(mid + 1, r)
            return root

        return helper(0, len(inorder) - 1)
```
