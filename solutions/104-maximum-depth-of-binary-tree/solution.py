from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def max_depth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        left_depth = self.max_depth(root.left)
        right_depth = self.max_depth(root.right)
        return max(left_depth, right_depth) + 1


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([1, 2, 3, 4, 5, 6, 7])
    got = sol.max_depth(root)
    status = "PASS" if got == 3 else "FAIL"
    print(f"{status} | depth={got}")
