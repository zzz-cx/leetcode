# 思路：递归，先遍历左子树，再遍历根节点，再遍历右子树
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def inorder_traversal(self, root: Optional[TreeNode]) -> List[int]:
        if root is None:
            return []
        out = self.inorder_traversal(root.left)
        out.append(root.val)
        right = self.inorder_traversal(root.right)
        out.extend(right)
        return out


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([1, 2, 3, 4, 5, 6, 7])
    got = sol.inorder_traversal(root)
    status = "PASS" if got == [4, 2, 5, 1, 6, 3, 7] else "FAIL"
    print(f"{status} | inorder={got}")
