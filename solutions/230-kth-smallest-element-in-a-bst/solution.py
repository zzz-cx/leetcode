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


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([3, 1, 4, None, 2])
    got = sol.kth_smallest(root, 1)
    status = "PASS" if got == 1 else "FAIL"
    print(f"{status} | kth={got}")
