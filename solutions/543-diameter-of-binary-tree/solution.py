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


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([1, 2, 3, 4, 5, 6, 7])
    got = sol.diameter_of_binary_tree(root)
    status = "PASS" if got == 6 else "FAIL"
    print(f"{status} | diameter={got}")
