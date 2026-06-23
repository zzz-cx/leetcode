from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def invert_tree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if root is None:
            return None
        left = self.invert_tree(root.left)
        right = self.invert_tree(root.right)
        root.left = right
        root.right = left
        return root


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([1, 2, 3, 4, 5, 6, 7])
    got = sol.invert_tree(root)
    status = "PASS" if got and got.val == 1 else "FAIL"
    print(f"{status} | inverted root={got.val if got else None}")
