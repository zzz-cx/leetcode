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


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([1, 2, 2, 3, 4, 4, 3])
    got = sol.is_symmetric(root)
    status = "PASS" if got is True else "FAIL"
    print(f"{status} | symmetric={got}")
