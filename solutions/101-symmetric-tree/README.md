# 对称二叉树

> LeetCode 101 · [symmetric-tree](https://leetcode.cn/problems/symmetric-tree/)

## 题目

判断二叉树是否轴对称。

## 题解思路与解析

- 思路：递归，判断左右子树是否对称
- 判断左右子树是否对称：如果左右子树都为空，则对称；如果左右子树有一个为空，则不对称；如果左右子树的值不相等，则不对称；如果左右子树的值相等，则继续判断左右子树的左右子树是否对称

## 解答

### Golang

```go
// 思路：递归，判断左右子树是否对称
// 判断左右子树是否对称：如果左右子树都为空，则对称；如果左右子树有一个为空，则不对称；如果左右子树的值不相等，则不对称；如果左右子树的值相等，则继续判断左右子树的左右子树是否对称
func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}
```

### Python

```python
# 思路：递归，判断左右子树是否对称
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def is_symmetric(self, root: Optional[TreeNode]) -> bool:
        return self._is_mirror(root.left, root.right)

    def _is_mirror(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if p is None and q is None:
            return True
        if p is None or q is None:
            return False
        return (
            p.val == q.val
            and self._is_mirror(p.left, q.right)
            and self._is_mirror(p.right, q.left)
        )
```
