from typing import Optional


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def lowest_common_ancestor(
        self, root: TreeNode, p: TreeNode, q: TreeNode
    ) -> Optional[TreeNode]:
        if root is None or root == p or root == q:
            return root
        left = self.lowest_common_ancestor(root.left, p, q)
        right = self.lowest_common_ancestor(root.right, p, q)
        if left is not None and right is not None:
            return root
        if left is not None:
            return left
        return right


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([3, 5, 1, 6, 2, 0, 8, None, None, 7, 4])
    # find nodes by value
    def find(node, v):
        if not node:
            return None
        if node.val == v:
            return node
        return find(node.left, v) or find(node.right, v)
    p, q = find(root, 5), find(root, 1)
    got = sol.lowest_common_ancestor(root, p, q)
    status = "PASS" if got and got.val == 3 else "FAIL"
    print(f"{status} | lca={got.val if got else None}")
