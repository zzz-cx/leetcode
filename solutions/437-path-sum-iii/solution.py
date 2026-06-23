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


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([1, 2, 3, 4, 5, 6, 7])
    got = sol.path_sum(root, 3)
    status = "PASS" if got == 2 else "FAIL"
    print(f"{status} | paths={got}")
